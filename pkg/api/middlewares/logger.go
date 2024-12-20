package middlewares

import (
	"go-boilerplate-app/pkg/utils/constants"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoggerMiddleware() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "${time_custom} ${user_agent} ${remote_ip} ${status} ${method} ${uri} ${latency_human} ${bytes_out}\n",
		CustomTimeFormat: constants.DEFAULT_LOGGER_TIMESTAMP_FORMAT,
	})
}
