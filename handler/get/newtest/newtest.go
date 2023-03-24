
package newtest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Newtest(c echo.Context) error {
	var newtest = "newtest"
	return c.Render(http.StatusOK, "newtest.html", map[string]interface{}{
		"newtest":newtest,
	})

}