package main

import (
	"context"
	"log"
	"os"

	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/pkg/database"
	"gitlab.com/chalermphanEur/server"
)

func main() {
	ctx := context.Background()
	_ = ctx

	//LoadConfig()
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Please provide config file path")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {

	case "kbm":
		dbMongo := database.DbConn(ctx, &cfg)

		defer dbMongo.Disconnect(ctx)

		server.StartMongo(ctx, &cfg, dbMongo)

	case "auth":
		dbSql := database.DbSqlServer(ctx, &cfg)

		server.StartSqlServer(ctx, &cfg, dbSql)

	case "users":
		dbSql := database.DbSqlServer(ctx, &cfg)

		server.StartSqlServer(ctx, &cfg, dbSql)

	case "elastic":

	}
	//Now don't use mongodb

}
