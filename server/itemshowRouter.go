package server

import (

	
	_itemshopControllor "github.com/neth/isekai-shop/pkg/itemShop/controller"
	_itemShoprepository "github.com/neth/isekai-shop/pkg/itemShop/repository"
	_itemShopService "github.com/neth/isekai-shop/pkg/itemShop/service_usecase"
)


func (s *echoServer) initItemshopRouter() {
	
	router := s.app.Group("/v1/item-shop")

	itemshoprepository := _itemShoprepository.NewitemShowrepositoryIpml(s.db, s.app.Logger)
	itemshopcontrollor := _itemShopService.NewitemShopServiceImpl(itemshoprepository)
	itemshopService := _itemshopControllor.NewitemShopControllerImpl(itemshopcontrollor)
	router.GET("", itemshopService.Listing)
} 

