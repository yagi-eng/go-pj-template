package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/yagi-eng/go-pj-template/controllers"
	"gorm.io/gorm"
)

// Init ルーティング設定
func Init(e *echo.Echo, db *gorm.DB) {
	api := e.Group("/api")
	{
		users := api.Group("/users")
		{
			users.GET("", controllers.GetUsers(db))
			users.GET("/add", controllers.CreateUser(db))
		}
	}
}
