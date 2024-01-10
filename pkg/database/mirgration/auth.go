package mirgration

import (
	"context"

	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/auth"
	"gitlab.com/chalermphanEur/pkg/database"
	"gorm.io/gorm"
)

func authDbcon(ctx context.Context, cfg *config.Config) *gorm.DB {
	return database.DbSqlServer(ctx, cfg)
}

func AuthMigration(ctx context.Context, cfg *config.Config) {
	db := authDbcon(ctx, cfg)

	db.AutoMigrate(&auth.Auth{})

	if err := db.Error; err != nil {
		panic(err)
	}
}
