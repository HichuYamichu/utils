package handler

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"net/http"

	"github.com/hichuyamichu-me/utils/app/store"
	imageService "github.com/hichuyamichu-me/utils/services/image"
	"github.com/labstack/echo"
)

type ImageResponce struct {
	Path string
}

type imageServiceFunc func(img *image.Image, a *imageService.Args) (*image.NRGBA, error)

func ForImageService(s *store.Store, h imageServiceFunc) echo.HandlerFunc {
	return echo.HandlerFunc(func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Input file missing.")
		}

		src, err := file.Open()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Input file malformed.")
		}
		defer src.Close()

		img, _, err := image.Decode(src)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Input file is not an image.")
		}

		args := &imageService.Args{}
		args.X = c.FormValue("x")
		args.Y = c.FormValue("y")
		args.Filter = c.FormValue("filter")
		args.Value = c.FormValue("value")
		res, err := h(&img, args)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		buffer := new(bytes.Buffer)
		if err := jpeg.Encode(buffer, res, nil); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not encode the file.")
		}

		fName, err := s.FS.SaveTemp(buffer.Bytes(), "image_service_")
		if err := jpeg.Encode(buffer, res, nil); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Could not save the file.")
		}

		fPath := fmt.Sprintf("http://localhost:3000/api/files/%s", fName)

		r := &ImageResponce{Path: fPath}
		return c.JSON(200, r)
	})
}
