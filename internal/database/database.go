package database

import (
	"ip-api/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection using GORM
func InitDB(dsn string) (db *gorm.DB, err error) {

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, db.AutoMigrate(&models.IP{})
}
