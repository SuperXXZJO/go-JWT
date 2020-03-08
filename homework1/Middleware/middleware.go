package Middleware

import (
	"github.com/labstack/echo"
	"homework1/User"
)

func Check (next echo.HandlerFunc)echo.HandlerFunc{
	return func (c echo.Context)error{
		token:= c.Param("token")
		err := User.Checktoken(token)
		if err != nil{
			return err
		}
		return next(c)
	}
}
