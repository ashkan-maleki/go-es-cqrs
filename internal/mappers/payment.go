package mappers

import (
	"github.com/ashkan-maleki/go-es-cqrs/internal/dto"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/models"
	orderService "github.com/ashkan-maleki/go-es-cqrs/proto/order"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func PaymentFromProto(protoPayment *orderService.Payment) dto.Payment {
	return dto.Payment{
		PaymentID: protoPayment.GetID(),
		Timestamp: protoPayment.GetTimestamp().AsTime(),
	}
}

func PaymentResponseFromModel(payment models.Payment) dto.Payment {
	return dto.Payment{
		PaymentID: payment.PaymentID,
		Timestamp: payment.Timestamp,
	}
}

func PaymentToProto(payment dto.Payment) *orderService.Payment {
	return &orderService.Payment{
		ID:        payment.PaymentID,
		Timestamp: timestamppb.New(payment.Timestamp),
	}
}
