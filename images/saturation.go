package images

import (
	"image"
	"net/url"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Saturation(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	saturation, err := strconv.ParseFloat(form.Get("saturation"), 64)
	if err != nil {
		return nil, appErrors.New(400, "saturation must be a float")
	}
	return imaging.AdjustSaturation(*img, saturation), nil
}
