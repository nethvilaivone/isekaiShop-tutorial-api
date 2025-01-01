package databases

import (
	"fmt"
	"log"
	"sync"

	"github.com/neth/isekai-shop/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
    *gorm.DB
}

func (db *postgresDatabase) ConnectedGetting() *gorm.DB {
    return db.DB
}

var (
    postgresDatabaseInstace *postgresDatabase
    once                    sync.Once
)

func NewPostgresDatabase(conf *config.Database) Databases {
    once.Do(func() {
        dsn := fmt.Sprintf(
            "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
            conf.Host,
            conf.Port,
            conf.User,
            conf.Password,    // Changed %d to %s since password is now a string
            conf.DBName,
            conf.SSLMode,
            conf.Schema,      // Changed schma to search_path (correct Postgres terminology)
        )

        conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
        if err != nil {
            panic(err)
        }

        log.Printf("Connected to database %s", conf.DBName)

        postgresDatabaseInstace = &postgresDatabase{conn}
    })

    return postgresDatabaseInstace
}