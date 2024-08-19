package mailsync

import (
	"context"
	"fmt"
	"log/slog"
	"sync"
	"time"

	"github.com/mxk/go-imap/imap"
	"github.com/tierklinik-dobersberg/mailbox"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// State is the state of a mail syncer.
type State struct {
	// ID is the MongoDB object ID.
	ID primitive.ObjectID `bson:"_id,omitempty"`
	// Name is the name of the mail syncer the
	// state belongs to.
	Name string `bson:"name,omitempty"`
	// UIDValidtity is the last UIDVALIDITY value
	// seen.
	UIDValidity uint32 `bson:"uidValidity,omitempty"`
	// LastUIDFetched is the last mail UID that has
	// been fetched.
	// Only valid if UIDValidity is still unchanged.
	LastUIDFetched uint32 `bson:"lastUidFetched,omitempty"`
}

// MessageHandler is called for each new E-Mail that has been received
// by the watched mailbox.
type MessageHandler interface {
	HandleMail(ctx context.Context, mail *mailbox.EMail)
}

// MessageHandlerFunc is a convenience type for implementing
// MessageHandler.
type MessageHandlerFunc func(context.Context, *mailbox.EMail)

// HandleMail implements MessageHandler.
func (fn MessageHandlerFunc) HandleMail(ctx context.Context, mail *mailbox.EMail) {
	fn(ctx, mail)
}

// Syncer syncs mails with mongodb.
type Syncer struct {
	rw           sync.Mutex
	state        State
	close        chan struct{}
	syncState    Store
	cfg          *mailbox.Config
	handler      MessageHandler
	pollInterval time.Duration
	wg           sync.WaitGroup
}

// OnMessage configures the message handler.
func (sync *Syncer) OnMessage(handler MessageHandler) {
	sync.rw.Lock()
	defer sync.rw.Unlock()

	sync.handler = handler
}

// Start starts the syncer.
func (sync *Syncer) Start() error {
	sync.rw.Lock()
	defer sync.rw.Unlock()

	if sync.close != nil {
		return fmt.Errorf("already running")
	}

	sync.close = make(chan struct{})
	closeCh := sync.close

	sync.wg.Add(1)
	go func() {
		defer sync.wg.Done()
		defer func() {
			sync.rw.Lock()
			defer sync.rw.Unlock()

			sync.close = nil
		}()

		for {
			sync.poll()

			select {
			case <-closeCh:
				return
			case <-time.After(sync.pollInterval):
				slog.Info("starting update")
			}
		}
	}()

	return nil
}

func (sync *Syncer) poll() {
	defer func() {
		if x := recover(); x != nil {
			slog.Error("recovered panic", slog.Any("panic", x))
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	slog.InfoContext(ctx, "connecting to IMAP server ...")
	cli, err := mailbox.Connect(*sync.cfg)
	if err != nil {
		slog.ErrorContext(ctx, "failed to connect", slog.Any("error", err.Error()))

		return
	}
	defer func() {
		_, err := cli.IMAP.Logout(time.Second * 2)
		if err != nil {
			slog.ErrorContext(ctx, "failed to perform IMAP logout", slog.Any("error", err.Error()))
		}
	}()

	if cli.IMAP.Mailbox.UIDValidity != sync.state.UIDValidity {
		slog.InfoContext(ctx, "mailbox UID validity changed", slog.Any("old", sync.state.UIDValidity), slog.Any("new", cli.IMAP.Mailbox.UIDValidity))
		// If UIDValidity changed we cannot continue to sync using the last retrieved
		// UID as it might be invalid now. We MUST start from scratch and re-sync all mails.
		sync.state.LastUIDFetched = 1
		sync.state.UIDValidity = cli.IMAP.Mailbox.UIDValidity

		// TODO(ppacher): we should somehow notify the MessageHandler about
		// this situation.
	}

	seqset := new(imap.SeqSet)
	seqset.AddRange(sync.state.LastUIDFetched, 0)

	slog.InfoContext(ctx, "searching for new mails", slog.Any("seqset", seqset.String()))
	mails, err := cli.FetchUIDs(ctx, seqset)
	if err != nil {
		slog.ErrorContext(ctx, "failed to fetch mails", slog.Any("error", err.Error()))

		return
	}

	var highestUID uint32
	count := 0
	for mail := range mails {
		if mail.Err != nil {
			slog.ErrorContext(ctx, "failed to fetch mail", slog.Any("error", err.Error()))

			continue
		}

		if mail.UID <= sync.state.LastUIDFetched {
			// we've already processed this mail.
			continue
		}

		sync.handler.HandleMail(ctx, mail.EMail)
		if mail.UID > highestUID {
			highestUID = mail.UID
			sync.updateState(ctx, cli.IMAP.Mailbox.UIDValidity, highestUID)
		}
		count++
	}

	slog.InfoContext(ctx, "new mail processed successfully", slog.Any("count", count))
}

func (sync *Syncer) updateState(ctx context.Context, uidvalidtity uint32, lastUID uint32) {
	sync.state.LastUIDFetched = lastUID
	sync.state.UIDValidity = uidvalidtity

	if err := sync.syncState.SaveState(ctx, sync.state); err != nil {
		slog.ErrorContext(ctx, "failed to persist sync state", slog.Any("error", err.Error()))
	}
}

// Stop stops the syncer.
func (sync *Syncer) Stop() error {
	sync.rw.Lock()
	if sync.close == nil {
		sync.rw.Unlock()

		return fmt.Errorf("not running")
	}
	close(sync.close)
	sync.rw.Unlock()

	sync.wg.Wait()

	return nil
}
