package app

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	"github.com/hichuyamichu-me/utils/app/handler"
	"github.com/hichuyamichu-me/utils/app/store"
	"github.com/hichuyamichu-me/utils/services/image"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := store.New()

	api := e.Group("/api")
	api.Static("/files", "tmp")
	imageService := api.Group("/images")
	imageService.POST("/fit", handler.ForImageService(s, image.Fit))
	imageService.POST("/fill", handler.ForImageService(s, image.Fill))
	imageService.POST("/resize", handler.ForImageService(s, image.Resize))
	imageService.POST("/blurr", handler.ForImageService(s, image.Blurr))
	imageService.POST("/saturate", handler.ForImageService(s, image.Saturation))
	imageService.POST("/sharpen", handler.ForImageService(s, image.Sharpen))
	imageService.POST("/gamma", handler.ForImageService(s, image.Gamma))
	imageService.POST("/contrast", handler.ForImageService(s, image.Contrast))

	return e
}
