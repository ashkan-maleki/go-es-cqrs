package v1

import (
	"context"

	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/aggregate"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/es"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
	"github.com/opentracing/opentracing-go/log"
)

type PayOrderCommandHandler interface {
	Handle(ctx context.Context, command *PayOrderCommand) error
}

type payOrderCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewOrderPaidHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *payOrderCommandHandler {
	return &payOrderCommandHandler{log: log, cfg: cfg, es: es}
}

func (c *payOrderCommandHandler) Handle(ctx context.Context, command *PayOrderCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "payOrderCommandHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.PayOrder(ctx, command.Payment); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
