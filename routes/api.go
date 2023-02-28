package routes

import (
	"tiny-url/handlers"

	"github.com/labstack/echo"
)

func InitializeGroup(e *echo.Echo , handler *handlers.Handlers) {
	api:=e.Group("/api/v1")
	api.GET("/:shortUrl", handler.ResolvUrlHandler())
	api.POST("/shorten", handler.ShortenUrlHandler())
}