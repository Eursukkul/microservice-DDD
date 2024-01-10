package main

import (
	"context"
	"log"
	"os"

	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/pkg/database/mirgration"
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

	// import statements

	switch cfg.App.Name {
	case "auth":
		mirgration.AuthMigration(ctx, &cfg)

	case "users":
		mirgration.UserMigration(ctx, &cfg)

	}
}
