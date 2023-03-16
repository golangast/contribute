package routes

import (
	"github.com/golangast/contribute/handler/get/home"
	"github.com/labstack/echo/v4"
)

//Routes is for routing
func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
}
