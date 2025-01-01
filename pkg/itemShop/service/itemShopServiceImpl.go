package service

import (
	_itemShoprepository "github.com/neth/isekai-shop/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShoprepository.ItemShopRepository
}

func NewitemShopServiceImpl(itemShopRepository _itemShoprepository.ItemShopRepository) ItemShopService {

	return &itemShopServiceImpl{itemShopRepository}
}
