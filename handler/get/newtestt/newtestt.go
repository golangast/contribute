
package newtestt

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Newtestt(c echo.Context) error {
	var newtestt = "newtestt"
	return c.Render(http.StatusOK, "newtestt.html", map[string]interface{}{
		"newtestt":newtestt,
	})

}