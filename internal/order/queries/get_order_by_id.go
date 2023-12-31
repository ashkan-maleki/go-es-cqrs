package queries

import (
	"context"
	"github.com/ashkan-maleki/go-es-cqrs/config"
	"github.com/ashkan-maleki/go-es-cqrs/internal/mappers"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/aggregate"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/models"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/repository"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/es"
	"github.com/ashkan-maleki/go-es-cqrs/pkg/logger"
	"github.com/opentracing/opentracing-go/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type GetOrderByIDQueryHandler interface {
	Handle(ctx context.Context, command *GetOrderByIDQuery) (*models.OrderProjection, error)
}

type getOrderByIDHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        es.AggregateStore
	mongoRepo repository.OrderMongoRepository
}

func NewGetOrderByIDHandler(log logger.Logger, cfg *config.Config, es es.AggregateStore, mongoRepo repository.OrderMongoRepository) *getOrderByIDHandler {
	return &getOrderByIDHandler{log: log, cfg: cfg, es: es, mongoRepo: mongoRepo}
}

func (q *getOrderByIDHandler) Handle(ctx context.Context, query *GetOrderByIDQuery) (*models.OrderProjection, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "getOrderByIDHandler.Handle")
	defer span.Finish()
	span.LogFields(log.String("AggregateID", query.ID))

	orderProjection, err := q.mongoRepo.GetByID(ctx, query.ID)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}
	if orderProjection != nil {
		return orderProjection, nil
	}

	order := aggregate.NewOrderAggregateWithID(query.ID)
	if err := q.es.Load(ctx, order); err != nil {
		return nil, err
	}

	if aggregate.IsAggregateNotFound(order) {
		return nil, aggregate.ErrOrderNotFound
	}

	orderProjection = mappers.OrderProjectionFromAggregate(order)

	_, err = q.mongoRepo.Insert(ctx, orderProjection)
	if err != nil {
		return nil, err
	}

	return orderProjection, nil
}
