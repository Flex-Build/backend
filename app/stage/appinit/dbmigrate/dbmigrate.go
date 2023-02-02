// Package dbmigrate provides method to migrate models into database
package dbmigrate

import (
	"github.com/Flexi-Build/backend/models"
	"github.com/Flexi-Build/backend/pkg/store"
	"github.com/TheLazarusNetwork/go-helpers/logo"
)

func Migrate() {
	db := store.DB
	err := db.AutoMigrate(
		&models.Site{},
	)
	if err != nil {
		logo.Fatalf("failed to migrate models into database: %s", err)
	}
}
