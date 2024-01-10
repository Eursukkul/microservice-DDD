package mirgration

import (
	"context"

	"gitlab.com/chalermphanEur/config"
	"gitlab.com/chalermphanEur/modules/users"
	"gitlab.com/chalermphanEur/pkg/database"
	"gorm.io/gorm"
)

func userDbcon(ctx context.Context, cfg *config.Config) *gorm.DB {
	return database.DbSqlServer(ctx, cfg)
}

func UserMigration(ctx context.Context, cfg *config.Config) {
	db := userDbcon(ctx, cfg)

	db.AutoMigrate(&users.User{})

	if err := db.Error; err != nil {
		panic(err)
	}
}
