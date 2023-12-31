package mappers

import (
	"github.com/ashkan-maleki/go-es-cqrs/internal/dto"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/events/v1"
)

func CreateOrderDtoToEventData(createDto dto.CreateOrderReqDto) v1.OrderCreatedEvent {
	return v1.OrderCreatedEvent{
		ShopItems:       createDto.ShopItems,
		AccountEmail:    createDto.AccountEmail,
		DeliveryAddress: createDto.DeliveryAddress,
	}
}
