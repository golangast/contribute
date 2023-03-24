package routes

import (
	"github.com/golangast/contribute/handler/get/home"
	"github.com/golangast/contribute/handler/get/news"
	"github.com/golangast/contribute/handler/get/newtest"
	"github.com/golangast/contribute/handler/get/newtestt"
	"github.com/labstack/echo/v4"
	"github.com/golangast/contribute/handler/get/newtestta"
//#import
)

//Routes is for routing
func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)

	e.GET("/news", news.News)
	e.GET("/newtest", newtest.Newtest)
	e.GET("/newtestt", newtestt.Newtestt)
	e.GET("/newtestta", newtestta.Newtestta)
//#routes
}
