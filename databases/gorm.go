package databases

import (
	"fmt"
	"miniproject/configs"
	"miniproject/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDatabase() {
	username := configs.Cfg.DBUsername
	password := configs.Cfg.DBPassword
	addrress := configs.Cfg.DBAddress
	dbName := configs.Cfg.DBName

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, addrress, dbName)

	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		panic("Failed to Connect database")
	}

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Food{}, &models.Order{}, &models.OrderItem{}, &models.Payment{})
}
