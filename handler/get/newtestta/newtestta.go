
package newtestta

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Newtestta(c echo.Context) error {
	var newtestta = "newtestta"
	return c.Render(http.StatusOK, "newtestta.html", map[string]interface{}{
		"newtestta":newtestta,
	})

}