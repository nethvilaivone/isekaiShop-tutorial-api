package entitise

import (
	"time"

	_itemShopModel "github.com/neth/isekai-shop/pkg/itemShop/model"
)

type Item struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	AdminID     *string   `gorm:"type:varchar(64);"`
	Name        string    `gorm:"type:varchar(64);unique;not null;"`
	Description string    `gorm:"type:varchar(128);not null;"`
	Picture     string    `gorm:"type:varchar(256);not null;"`
	Price       uint      `gorm:"not null;"`
	IsArchive   bool      `gorm:"not null;default:false;"`
	CreatedAt   time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt   time.Time `gorm:"not null;autoUpdateTime;"`
}

// convert entitise to Model  
func (i *Item) ToitemModel() *_itemShopModel.Item {
	return &_itemShopModel.Item{
		ID:          i.ID,
		Name:        i.Name,
		Description: i.Description,
		Picture:     i.Picture,
		Price:       i.Price,
	}

}
