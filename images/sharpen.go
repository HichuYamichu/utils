package images

import (
	"image"
	"net/url"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Sharpen(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	sigma, err := strconv.ParseFloat(form.Get("sigma"), 64)
	if err != nil {
		return nil, appErrors.New(400, "sigma must be a float")
	}
	return imaging.Sharpen(*img, sigma), nil
}
