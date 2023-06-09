package fxpubsub

import (
	"context"
	"fmt"

	"cloud.google.com/go/pubsub"
	"github.com/ekkinox/fx-template/modules/fxconfig"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

var FxPubSubModule = fx.Module(
	"pubsub",
	fx.Provide(
		NewFxPubSub,
	),
)

type FxPubSubParam struct {
	fx.In
	LifeCycle      fx.Lifecycle
	Config         *fxconfig.Config
	TracerProvider *trace.TracerProvider
}

func NewFxPubSub(p FxPubSubParam) (*pubsub.Client, error) {

	// client
	client, err := pubsub.NewClient(context.Background(), p.Config.GetString("pubsub.project.id"))
	if err != nil {
		return nil, fmt.Errorf("failed to create pubsub client: %w", err)
	}

	// lifecycle
	p.LifeCycle.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
