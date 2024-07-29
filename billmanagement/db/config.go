package config

import (
	app "com.electricity.online.bill/config"
	"com.electricity.online.bill/models"

	"github.com/micro/micro/v3/service/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DatabaseClient *gorm.DB

func InitializeDb() {
	DatabaseClient = GetClient()
	if DatabaseClient == nil {
		logger.Error("Database client is nil")
	}
}

func GetClient() *gorm.DB {
	db_url := app.GetVal("GO_MICRO_DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})
	if err != nil {
		logger.Errorf("failed to connect database")
		return nil
	}
	db.AutoMigrate(&models.Bill{})
	return db
}
