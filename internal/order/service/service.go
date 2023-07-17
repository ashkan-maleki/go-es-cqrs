package service

import (
	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/commands/v1"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/queries"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/repository"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/es"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
)

type OrderService struct {
	Commands *v1.OrderCommands
	Queries  *queries.OrderQueries
}

func NewOrderService(
	log logger.Logger,
	cfg *config.Config,
	es es.AggregateStore,
	mongoRepo repository.OrderMongoRepository,
	elasticRepository repository.ElasticOrderRepository,
) *OrderService {

	createOrderHandler := v1.NewCreateOrderHandler(log, cfg, es)
	orderPaidHandler := v1.NewOrderPaidHandler(log, cfg, es)
	submitOrderHandler := v1.NewSubmitOrderHandler(log, cfg, es)
	updateOrderCmdHandler := v1.NewUpdateShoppingCartCmdHandler(log, cfg, es)
	cancelOrderCommandHandler := v1.NewCancelOrderCommandHandler(log, cfg, es)
	deliveryOrderCommandHandler := v1.NewCompleteOrderCommandHandler(log, cfg, es)
	changeOrderDeliveryAddressCmdHandler := v1.NewChangeDeliveryAddressCmdHandler(log, cfg, es)

	getOrderByIDHandler := queries.NewGetOrderByIDHandler(log, cfg, es, mongoRepo)
	searchOrdersHandler := queries.NewSearchOrdersHandler(log, cfg, es, elasticRepository)

	orderCommands := v1.NewOrderCommands(
		createOrderHandler,
		orderPaidHandler,
		submitOrderHandler,
		updateOrderCmdHandler,
		cancelOrderCommandHandler,
		deliveryOrderCommandHandler,
		changeOrderDeliveryAddressCmdHandler,
	)
	orderQueries := queries.NewOrderQueries(getOrderByIDHandler, searchOrdersHandler)

	return &OrderService{Commands: orderCommands, Queries: orderQueries}
}
