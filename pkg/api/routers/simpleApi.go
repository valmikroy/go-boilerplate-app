package router

// this is where middleware and handlers are registered

import (
	"go-boilerplate-app/pkg/api/handlers/errors"
	healthHandlers "go-boilerplate-app/pkg/api/handlers/healthz"
	"go-boilerplate-app/pkg/api/middlewares"
	"go-boilerplate-app/pkg/logger"
)

var simpleAPIRouter *Router

func InitSimpleAPIRouter() {

	logger.Info("Initializing Simple API router")

	simpleAPIRouter = &Router{}
	simpleAPIRouter.Name = "Simple API"
	simpleAPIRouter.Init()

	logger.Info("Registering simple API pre middlewares")
	simpleAPIRouter.RegisterPreMiddleware(middlewares.SlashesMiddleware())

	// register logging
	simpleAPIRouter.RegisterMiddleware(middlewares.LoggerMiddleware())
	// Request middleware
	simpleAPIRouter.RegisterMiddleware(middlewares.RequestHeadersMiddleware())
	// Response middleware
	simpleAPIRouter.RegisterMiddleware(middlewares.ResponseHeadersMiddleware())

	// healthcheck handlers
	registerSimpleApiHealthCheckHandlers()

	// error handlers
	registerSimpleApiErrorHandlers()

}

func SimpleAPIRouter() *Router {
	return simpleAPIRouter
}

func registerSimpleApiHealthCheckHandlers() {
	health := simpleAPIRouter.Echo.Group("/health")
	health.GET("/alive", healthHandlers.Index)
	health.GET("/ready", healthHandlers.Ready)
}

func registerSimpleApiErrorHandlers() {
	simpleAPIRouter.Echo.HTTPErrorHandler = errors.AutomatedHttpErrorHandler()
	simpleAPIRouter.Echo.RouteNotFound("/*", errors.NotFound)
}
