package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/yagi-eng/go-pj-template/infrastructure"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("Error loading env: %v", err)
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	// アクセス元を制限する場合は上記を消してこちらを使用
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"(アクセスを許可するURL)"},
	// }))

	// DB Connect
	db := infrastructure.Connect()

	// Routes
	infrastructure.Init(e, db)

	// Start server
	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
