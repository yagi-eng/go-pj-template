package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/yagi-eng/go-pj-template/infrastructure"
	"github.com/yagi-eng/go-pj-template/model"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading env: %v", err)
	}

	db := infrastructure.Connect()

	db.AutoMigrate(&model.User{})
}
