package test

import (
	"errors"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"go-boilerplate-app/pkg/api/handlers"
	"go-boilerplate-app/pkg/api/helpers"
	"go-boilerplate-app/pkg/utils/constants"

	"github.com/labstack/echo/v4"
)

func Get(c echo.Context) error {

	var intCheck = regexp.MustCompile(`^[0-9]+$`)
	var m int

	max := c.Param("max")

	if max == "" {
		m = 6
	} else if intCheck.MatchString(max) {
		var e error
		m, e = strconv.Atoi(max)
		if e != nil {
			return helpers.Error(c, e, nil)
		}
	} else {
		return helpers.Error(c, constants.ERROR_INVALID_DATA, errors.New("invalid input: please provide integer"))
	}

	r := 1 + rand.Intn(m)
	payload := map[string]string{
		"message": strconv.Itoa(r),
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))
}
