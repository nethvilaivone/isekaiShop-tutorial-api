package main

import (
	"github.com/neth/isekai-shop/config"
	"github.com/neth/isekai-shop/databases"
	"github.com/neth/isekai-shop/entitise"

)

func main() {

	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)

	// tx := db.ConnectedGetting().Begin()

	playerMigration(db)
	adminMigration(db)
	playCoinMigration(db)
	itemMigration(db)
	inventoryMigration(db)
	purchaseMigration(db)

	// if err := tx.Commit(); err != nil {
	// 	tx.Rollback()
	// 	panic(err)
	// }
}

func playerMigration(tx databases.Databases) {

	tx.ConnectedGetting().Migrator().CreateTable(&entitise.Player{})
}

func adminMigration(tx databases.Databases) {
	tx.ConnectedGetting().Migrator().CreateTable(entitise.Admin{})
}

func playCoinMigration(tx databases.Databases) {
	tx.ConnectedGetting().Migrator().CreateTable(&entitise.PlayerCoin{})
}

func itemMigration(tx databases.Databases) {
	tx.ConnectedGetting().Migrator().CreateTable(&entitise.Item{})

}

func inventoryMigration(tx databases.Databases) {
	tx.ConnectedGetting().Migrator().CreateTable(&entitise.Inventory{})
}

func purchaseMigration(tx databases.Databases) {
	tx.ConnectedGetting().Migrator().CreateTable(&entitise.PurchaseHistory{})
}
