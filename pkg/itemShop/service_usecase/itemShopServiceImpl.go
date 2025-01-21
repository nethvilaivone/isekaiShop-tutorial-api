package service

import (
	_itemShopModel "github.com/neth/isekai-shop/pkg/itemShop/model"
	_itemShoprepository "github.com/neth/isekai-shop/pkg/itemShop/repository"
)

type itemShopServiceImpl struct {
	itemShopRepository _itemShoprepository.ItemShopRepository
}

func NewitemShopServiceImpl(itemShopRepository _itemShoprepository.ItemShopRepository) ItemShopService {

	return &itemShopServiceImpl{itemShopRepository}
}

 
func (item *itemShopServiceImpl) Listing()([]*_itemShopModel.Item, error ){
 listing , err := 	item.itemShopRepository.Listing()

 if  err != nil  {
	return nil  , err 
 }


itmeModel :=  make([]*_itemShopModel.Item , 0)

for _, item := range listing {
	itmeModel = append(itmeModel, item.ToitemModel())
}

return itmeModel , nil

}
