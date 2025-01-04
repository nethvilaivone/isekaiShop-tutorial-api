package main

import (
	"github.com/neth/isekai-shop/config"
	"github.com/neth/isekai-shop/databases"
	"github.com/neth/isekai-shop/entitise"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	tx := db.ConnectedGetting().Begin()

	playerMigration(tx)
	adminMigration(tx)
	playCoinMigration(tx)
	itemMigration(tx)
	inventoryMigration(tx)
	purchaseMigration(tx)

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		panic(err)
	}
}


func playerMigration(tx *gorm.DB) {

	tx.Migrator().CreateTable(&entitise.Player{})
}

func adminMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(entitise.Admin{})
}

func playCoinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entitise.PlayerCoin{})
}

func itemMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entitise.Item{})

}

func inventoryMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entitise.Inventory{})
}

func purchaseMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entitise.PurchaseHistory{})
}
