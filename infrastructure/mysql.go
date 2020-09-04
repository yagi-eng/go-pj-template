package infrastructure

import (
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect DB接続
func Connect() (db *gorm.DB) {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_USERPASS") +
		"@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" +
		os.Getenv("DB_NAME") +
		"?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatalf("Error connect DB: %v", err)
	}

	return db
}
