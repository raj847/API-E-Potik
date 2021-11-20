package routes

import (
	"minpro_arya/app/controllers/admin"

	"github.com/labstack/echo/v4"
)

type HandlerRoute struct {
	AdminController admin.UserController
}

func (handler *HandlerRoute) RouteRegister(e *echo.Echo) {
	//Create Account
	creates := e.Group("create")
	//admin
	creates.POST("/register", handler.AdminController.Register)
}
