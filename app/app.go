package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hichuyamichu-me/utils/app/handler"
	"github.com/hichuyamichu-me/utils/services/image"
)

// New creates new app instance
func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := handler.New()

	api := e.Group("/api")
	api.Static("/files", "tmp")
	imageService := api.Group("/images")
	imageService.POST("/fill", h.ForImageService(image.Fit))
	imageService.POST("/fill", h.ForImageService(image.Fill))
	imageService.POST("/resize", h.ForImageService(image.Resize))
	imageService.POST("/blurr", h.ForImageService(image.Blurr))
	imageService.POST("/saturate", h.ForImageService(image.Saturation))
	imageService.POST("/sharpen", h.ForImageService(image.Sharpen))
	imageService.POST("/gamma", h.ForImageService(image.Gamma))
	imageService.POST("/contrast", h.ForImageService(image.Contrast))

	return e
}
