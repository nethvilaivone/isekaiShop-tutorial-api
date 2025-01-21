package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	 custom "github.com/neth/isekai-shop/pkg/custom"
	_ "github.com/neth/isekai-shop/pkg/itemShop/exception"
	_itemShopService "github.com/neth/isekai-shop/pkg/itemShop/service_usecase"
	_itemShopModel "github.com/neth/isekai-shop/pkg/itemShop/model"
)

type itemShopControllerImpl struct {
	itemShopservice _itemShopService.ItemShopService
}

func NewitemShopControllerImpl(itemShopservice _itemShopService.ItemShopService) ItemShopController {
	return &itemShopControllerImpl{itemShopservice: itemShopservice}
}

func (c *itemShopControllerImpl) Listing(pctx echo.Context) error {
	itemtemfiller := new(_itemShopModel.ItemFillter)
	customEchoRest := custom.NewCustomEchoResquest(pctx)

	if err := customEchoRest.Bind(itemtemfiller); err != nil {
		return custom.CustomErrorMessage(pctx, http.StatusBadRequest, err.Error())
	}

	itemshoplisting, err := c.itemShopservice.Listing()

	if err != nil {
		return custom.CustomErrorMessage(pctx, http.StatusInternalServerError, err.Error())
	}

	return pctx.JSON(http.StatusOK, itemshoplisting)
}

