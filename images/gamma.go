package images

import (
	"image"
	"net/url"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Gamma(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	gamma, err := strconv.ParseFloat(form.Get("gamma"), 64)
	if err != nil {
		return nil, appErrors.New(400, "gamma must be a float")
	}
	return imaging.AdjustGamma(*img, gamma), nil
}
