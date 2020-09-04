package controllers

import (
	"github.com/labstack/echo"
	"github.com/valyala/fasthttp"
	"github.com/yagi-eng/go-pj-template/model"
	"gorm.io/gorm"
)

// GetUsers ユーザを取得
func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []model.User
		db.Find(&users)
		return c.JSON(fasthttp.StatusOK, users)
	}
}

// CreateUser ユーザ追加
func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		userName := c.QueryParam("name")
		user := model.User{Name: userName}
		db.Create(&user)
		return c.JSON(fasthttp.StatusOK, user)
	}
}
