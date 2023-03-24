
package news

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func News(c echo.Context) error {
	var news = "news"
	return c.Render(http.StatusOK, "news.html", map[string]interface{}{
		"routes":news,
	})

}