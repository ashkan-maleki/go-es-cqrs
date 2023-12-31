package v1

import (
	"context"

	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/aggregate"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/es"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
	"github.com/opentracing/opentracing-go/log"
)

type UpdateShoppingCartCommandHandler interface {
	Handle(ctx context.Context, command *UpdateShoppingCartCommand) error
}

type updateShoppingCartCmdHandler struct {
	log logger.Logger
	cfg *config.Config
	es  es.AggregateStore
}

func NewUpdateShoppingCartCmdHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore) *updateShoppingCartCmdHandler {
	return &updateShoppingCartCmdHandler{log: log, cfg: cfg, es: es}
}

func (c *updateShoppingCartCmdHandler) Handle(ctx context.Context, command *UpdateShoppingCartCommand) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "updateShoppingCartCmdHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", command.GetAggregateID()))

	order, err := aggregate.LoadOrderAggregate(ctx, c.es, command.GetAggregateID())
	if err != nil {
		return err
	}

	if err := order.UpdateShoppingCart(ctx, command.ShopItems); err != nil {
		return err
	}

	return c.es.Save(ctx, order)
}
