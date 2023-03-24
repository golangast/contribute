package routes

var Routetemp = `
package {{.routes}}

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func {{title .routes}}(c echo.Context) error {
	var {{.routes}} = "{{.routes}}"
	return c.Render(http.StatusOK, "{{.routes}}.html", map[string]interface{}{
		"{{.routes}}":{{.routes}},
	})

}`
