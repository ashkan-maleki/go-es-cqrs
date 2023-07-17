package mappers

import (
	"github.com/ashkan-maleki/go-es-cqrs/internal/dto"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/events/v1"
)

func ChangeDeliveryAddressReqDtoToEventData(reqDto dto.ChangeDeliveryAddressReqDto) v1.OrderDeliveryAddressChangedEvent {
	return v1.OrderDeliveryAddressChangedEvent{
		DeliveryAddress: reqDto.DeliveryAddress,
	}
}
