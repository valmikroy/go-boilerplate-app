package healthz

import (
	"net/http"

	"go-boilerplate-app/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func Ready(c echo.Context) error {
	payload := map[string]string{
		"message": "ready",
	}
	return c.JSON(http.StatusOK, handlers.Success(payload))
}
