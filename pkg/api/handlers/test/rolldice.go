package test

import (
	"math/rand"
	"net/http"
	"strconv"

	"go-boilerplate-app/pkg/api/handlers"

	"github.com/labstack/echo/v4"
)

func Rolldice(c echo.Context) error {
	roll := 1 + rand.Intn(6) //nolint:gosec // G404: Use of weak random number generator (math/rand instead of crypto/rand) is ignored as this is not security-sensitive.
	payload := map[string]string{
		"message": strconv.Itoa(roll),
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))
}
