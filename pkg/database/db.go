package database

import (
	"context"
	"log"
	"time"

	"gitlab.com/chalermphanEur/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func DbSqlServer(pctx context.Context, cfg *config.Config) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	db, err := gorm.Open(sqlserver.Open(cfg.Db.Url), &gorm.Config{})

	if err != nil {
		log.Fatal("Connection Failed to Open")
	}

	sqlserverDb, err := db.DB()
	err = sqlserverDb.PingContext(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Db.Url))

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		log.Fatal(err)
	}

	return client
}
