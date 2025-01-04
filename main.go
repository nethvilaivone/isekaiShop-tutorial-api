package main

import (
	"github.com/neth/isekai-shop/config"
	"github.com/neth/isekai-shop/databases"
	"github.com/neth/isekai-shop/server"
)


func main(){
	conf := config.ConfigGetting()
	db := databases.NewPostgresDatabase(conf.Database)
	server:= server.NewechoServer(conf, db.ConnectedGetting())
	server.Start()
}