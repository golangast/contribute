
package router

import (
	"github.com/labstack/echo/v4"
	"contribute/handler/get/home"

//#import 
)

//Routes is for routing
func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)


//#routes
}
