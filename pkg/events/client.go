package events

import (
	"context"
	"errors"
	"io"
	"log/slog"
	"strings"
	"sync"
	"time"

	"github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"golang.org/x/exp/maps"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
)

var (
	ErrNotConnected = errors.New("not yet connected")
)

type stream = *connect.BidiStreamForClient[eventsv1.SubscribeRequest, eventsv1.Event]

type Client struct {
	lock         sync.Mutex
	eventsClient stream
	events       map[protoreflect.FullName][]chan<- *eventsv1.Event

	newSubscription chan protoreflect.FullName
	incoming        chan *eventsv1.Event
	httpClient      connect.HTTPClient
	startupComplete chan struct{}

	wg       sync.WaitGroup
	endpoint string
}

func NewClient(ep string, httpClient connect.HTTPClient) *Client {
	cli := &Client{
		events:          make(map[protoreflect.FullName][]chan<- *eventsv1.Event),
		endpoint:        ep,
		httpClient:      httpClient,
		newSubscription: make(chan protoreflect.FullName, 10),
		incoming:        make(chan *eventsv1.Event, 10),
		startupComplete: make(chan struct{}),
	}

	return cli
}

func (cli *Client) Start(ctx context.Context) error {
	ready := cli.startupComplete

	if ready == nil {
		return errors.New("already started")
	}

	cli.wg.Add(2)
	go cli.receiveLoop(ctx)
	go cli.handleMessages(ctx)

	<-ready

	return nil
}

func (cli *Client) SubscribeMessage(ctx context.Context, msg proto.Message) (<-chan *eventsv1.Event, error) {
	typeUrl := proto.MessageName(msg)

	return cli.Subscribe(ctx, typeUrl)
}

func (cli *Client) Subscribe(ctx context.Context, typeUrl protoreflect.FullName) (<-chan *eventsv1.Event, error) {
	result := make(chan *eventsv1.Event, 10)

	cli.lock.Lock()
	cli.events[typeUrl] = append(cli.events[typeUrl], result)
	cli.lock.Unlock()

	cli.newSubscription <- typeUrl

	return result, nil
}

func (cli *Client) Publish(ctx context.Context, msg proto.Message) error {
	eventsClient := eventsv1connect.NewEventServiceClient(cli.httpClient, cli.endpoint)
	pb, err := anypb.New(msg)
	if err != nil {
		return err
	}

	slog.Info("publishing event", slog.Any("typeUrl", pb.TypeUrl))

	_, err = eventsClient.Publish(ctx, connect.NewRequest(&eventsv1.Event{
		Event: pb,
	}))

	return err
}

func (cli *Client) receiveLoop(ctx context.Context) {
	defer cli.wg.Done()

	for {
		if ctx.Err() != nil {
			return
		}

		slog.Info("connecting to events service", slog.Any("endpoint", cli.endpoint))
		eventsClient := eventsv1connect.NewEventServiceClient(cli.httpClient, cli.endpoint)

		stream := eventsClient.Subscribe(ctx)

		if cli.startupComplete != nil {
			close(cli.startupComplete)
			cli.startupComplete = nil
		}

		cli.lock.Lock()
		cli.eventsClient = stream
		sub := maps.Keys(cli.events)
		cli.lock.Unlock()

		for _, s := range sub {
			cli.newSubscription <- s
		}

		for {
			msg, err := stream.Receive()
			if err != nil {
				if !errors.Is(err, io.EOF) {
					slog.Error("receive side closed unexpectedly", slog.Any("error", err.Error()))
				}
				break
			}

			cli.incoming <- msg
		}

		time.Sleep(time.Second)
	}
}

func (cli *Client) handleMessages(ctx context.Context) {
	defer cli.wg.Done()

	for {
		select {
		case sub := <-cli.newSubscription:
			cli.lock.Lock()
			eventsCli := cli.eventsClient
			cli.lock.Unlock()

			err := eventsCli.Send(&eventsv1.SubscribeRequest{
				Kind: &eventsv1.SubscribeRequest_Subscribe{
					Subscribe: string(sub),
				},
			})

			if err != nil {
				slog.Error("failed to send subscription request", slog.Any("error", err.Error()), slog.Any("typeUrl", sub))
			}

		case msg := <-cli.incoming:

			typeUrl := trimTypeUrl(msg.Event.TypeUrl)

			cli.lock.Lock()
			subscribers := cli.events[typeUrl]
			cli.lock.Unlock()

			if len(subscribers) == 0 {
				slog.Warn("no subscribers found for incoming message", slog.Any("typeUrl", typeUrl))
			}

			for _, ch := range subscribers {
				ch <- proto.Clone(msg).(*eventsv1.Event)
			}

		case <-ctx.Done():
			return
		}
	}
}

func trimTypeUrl(url string) protoreflect.FullName {
	return protoreflect.FullName(strings.TrimPrefix(url, "type.googleapis.com/"))
}
