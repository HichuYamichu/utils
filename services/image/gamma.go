package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Gamma(img *image.Image, opts options) (*image.NRGBA, error) {
	gamma, err := strconv.ParseFloat(opts.Get("gamma"), 64)
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("gamma value must be a float"))
	}
	return imaging.AdjustGamma(*img, gamma), nil
}
