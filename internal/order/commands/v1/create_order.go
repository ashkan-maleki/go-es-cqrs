package v1

import (
	"context"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/aggregate"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/es"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
)

type CreateOrderCommandHandler interface {
	Handle(ctx context.Context, command *CreateOrderCommand) error
}

type createOrderHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewCreateOrderHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *createOrderHandler {
	return &createOrderHandler{log: log, cfg: cfg, es: es}
}

func (c *createOrderHandler) Handle(ctx context.Context, command *CreateOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "createOrderHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order := aggregate.NewOrderAggregateWithID(command.AggregateID)
	err := c.es.Exists(ctx, order.GetID())
	if err != nil && !errors.Is(err, esdb.ErrStreamNotFound) {
		return err
	}

	if err := order.CreateOrder(ctx, command.ShopItems, command.AccountEmail, command.DeliveryAddress); err != nil {
		return err
	}

	span.LogFields(log.String("order", order.String()))
	return c.es.Save(ctx, order)
}
