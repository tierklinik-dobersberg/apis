package events

import (
	"context"
	"fmt"
	"time"

	"github.com/bufbuild/connect-go"
	eventsv1 "github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1"
	"github.com/tierklinik-dobersberg/apis/gen/go/tkd/events/v1/eventsv1connect"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/consuldiscover"
	"github.com/tierklinik-dobersberg/apis/pkg/discovery/wellknown"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func PublishTo(ctx context.Context, svc eventsv1connect.EventServiceClient, msg proto.Message, retained bool) error {
	anyMsg, err := anypb.New(msg)
	if err != nil {
		return fmt.Errorf("failed to create google.protobuf.Any: %w", err)
	}

	evt := &eventsv1.Event{
		Event:    anyMsg,
		Retained: retained,
	}

	req := connect.NewRequest(evt)
	wait := time.Second
	tries := 0

	for {
		_, err := svc.Publish(ctx, req)
		if err == nil {
			return nil
		}

		if tries > 10 {
			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		select {
		case <-time.After(wait):
			wait = wait * 2
			tries++

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func PublishWithConsul(ctx context.Context, msg proto.Message, retained bool) error {
	disc, err := consuldiscover.NewFromEnv()
	if err != nil {
		return fmt.Errorf("failed to get discovery client: %w", err)
	}

	svc, err := wellknown.EventService.Create(ctx, disc)
	if err != nil {
		return fmt.Errorf("failed to get events service client: %w", err)
	}

	return PublishTo(ctx, svc, msg, retained)
}
