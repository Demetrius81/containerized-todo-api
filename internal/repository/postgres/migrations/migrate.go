package migrations

import (
	"github.com/Demetrius81/containerized-todo-api/internal/domain"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	if !db.Migrator().HasTable(&domain.Todo{}) {
		if err := db.AutoMigrate(&domain.Todo{}); err != nil {
			return err
		}
	}

	return nil
}
