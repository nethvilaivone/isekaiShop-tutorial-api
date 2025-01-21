package repository

import "github.com/neth/isekai-shop/entitise"

type ItemShopRepository interface {
		Listing() ([]*entitise.Item, error)	
		

}