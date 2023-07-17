package dto

import "github.com/ashkan-maleki/go-es-cqrs/internal/order/models"

type UpdateShoppingItemsReqDto struct {
	ShopItems []*models.ShopItem `json:"shopItems" bson:"shopItems,omitempty" validate:"required"`
}
