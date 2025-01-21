package service

import (

	_itemShopModel "github.com/neth/isekai-shop/pkg/itemShop/model"
)

type ItemShopService interface {
	Listing() ([]*_itemShopModel.Item,error)
}