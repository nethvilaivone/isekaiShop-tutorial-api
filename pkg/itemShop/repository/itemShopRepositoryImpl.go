package repository

import (
	"github.com/labstack/echo/v4"
	"github.com/neth/isekai-shop/entitise"
	"gorm.io/gorm"
	_itemshopExection "github.com/neth/isekai-shop/pkg/itemShop/exception"

)

type itemShowrepositoryIpml struct {
	db     *gorm.DB
	loggor echo.Logger
}

func NewitemShowrepositoryIpml(db *gorm.DB, logger echo.Logger) ItemShopRepository {
	return &itemShowrepositoryIpml{db, logger}
}

func (i *itemShowrepositoryIpml) Listing() ([]*entitise.Item, error) {

itemsList := make([]*entitise.Item, 0)

if	err := i.db.Find(&itemsList).Error; err != nil {
  i.loggor.Errorf("Fail to listing items %s", err.Error())
  return nil ,&_itemshopExection.ItemListing{}
}

	return itemsList, nil
}


