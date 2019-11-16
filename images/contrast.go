package images

import (
	"image"
	"net/url"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Contrast(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	contrast, err := strconv.ParseFloat(form.Get("contrast"), 64)
	if err != nil {
		return nil, appErrors.New(400, "contrast must be a float")
	}
	return imaging.AdjustContrast(*img, contrast), nil
}
