package mappers

import (
	"github.com/ashkan-maleki/go-es-cqrs/internal/dto"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/events/v1"
)

func UpdateOrderReqDtoToEventData(reqDto dto.UpdateShoppingItemsReqDto) v1.ShoppingCartUpdatedEvent {
	return v1.ShoppingCartUpdatedEvent{
		ShopItems: reqDto.ShopItems,
	}
}
