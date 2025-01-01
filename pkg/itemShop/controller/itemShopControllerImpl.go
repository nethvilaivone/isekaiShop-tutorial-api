package controller 

import (
	_itemShopService "github.com/neth/isekai-shop/pkg/itemShop/service"
)

type itemShopControllerImpl struct {
	itemShopservice  _itemShopService.ItemShopService
}


func NewitemShopControllerImpl(itemShopservice _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopservice: itemShopservice}
}
