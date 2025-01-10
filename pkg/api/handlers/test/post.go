package test

import (
	"errors"
	"go-boilerplate-app/pkg/api/handlers"
	"go-boilerplate-app/pkg/api/helpers"
	"go-boilerplate-app/pkg/utils/constants"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Input struct {
	Min string `json:"min" validate:"required,min=1,max=5"`
	Max string `json:"max" validate:"required,min=1,max=5"`
}

func Post(c echo.Context) error {

	var intCheck = regexp.MustCompile(`^[0-9]{1,5}$`)
	var min, max int

	f := &Input{}

	if err := c.Bind(f); err != nil {
		return helpers.Error(c, constants.ERROR_BINDING_BODY, err)
	}

	if !intCheck.MatchString(f.Min) {
		return helpers.Error(c, constants.ERROR_INVALID_DATA, errors.New("invalid input: please provide Min integer"))
	}
	if !intCheck.MatchString(f.Max) {
		return helpers.Error(c, constants.ERROR_INVALID_DATA, errors.New("invalid input: please provide Max integer"))
	}

	min, _ = strconv.Atoi(f.Min)
	max, _ = strconv.Atoi(f.Max)

	if min > max {
		return helpers.Error(c, constants.ERROR_INVALID_DATA, errors.New("invalid input: please provide Max integer greater than Min"))
	}

	r := min + rand.Intn(max-min)
	payload := map[string]string{
		"message": strconv.Itoa(r),
	}

	return c.JSON(http.StatusOK, handlers.Success(payload))
}
