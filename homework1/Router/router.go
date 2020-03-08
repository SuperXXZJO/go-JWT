package Router

import (
	"github.com/labstack/echo"
	"homework1/Middleware"
	"homework1/User"
)

func RUN () {
	homework := echo.New()
	homework.POST("/signup",User.Signup)
	homework.POST("/login",User.Login)
	homework.GET("/Find/:username",User.FindUser)//查询用户信息
	homework.POST("/Update",User.UpdateUser,Middleware.Check)//修改用户信息
	homework.Start(":8080")

}