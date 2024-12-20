package healthz

import (
	"net/http"

	"go-boilerplate-app/pkg/api/handlers"
	"go-boilerplate-app/pkg/config"

	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) error {
	payload := map[string]string{
		"message": "ok",
		"version": config.Version,
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))
}
