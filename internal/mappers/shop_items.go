package mappers

import (
	"github.com/ashkan-maleki/go-es-cqrs/internal/dto"
	"github.com/ashkan-maleki/go-es-cqrs/internal/order/models"
	orderService "github.com/ashkan-maleki/go-es-cqrs/proto/order"
)

func ShopItemResponseFromModel(item *models.ShopItem) dto.ShopItem {
	return dto.ShopItem{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}

func ShopItemsResponseFromModels(items []*models.ShopItem) []dto.ShopItem {
	shopItems := make([]dto.ShopItem, 0, len(items))
	for _, item := range items {
		shopItems = append(shopItems, ShopItemResponseFromModel(item))
	}
	return shopItems
}

func ShopItemResponseFromProto(item *orderService.ShopItem) dto.ShopItem {
	return dto.ShopItem{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}

func ShopItemsResponseFromProto(items []*orderService.ShopItem) []dto.ShopItem {
	shopItems := make([]dto.ShopItem, 0, len(items))
	for _, item := range items {
		shopItems = append(shopItems, ShopItemResponseFromProto(item))
	}
	return shopItems
}

func ShopItemResponseToProto(item dto.ShopItem) *orderService.ShopItem {
	return &orderService.ShopItem{
		ID:          item.ID,
		Title:       item.Title,
		Description: item.Description,
		Quantity:    item.Quantity,
		Price:       item.Price,
	}
}

func ShopItemsResponseToProto(items []dto.ShopItem) []*orderService.ShopItem {
	shopItems := make([]*orderService.ShopItem, 0, len(items))
	for _, item := range items {
		shopItems = append(shopItems, ShopItemResponseToProto(item))
	}
	return shopItems
}
