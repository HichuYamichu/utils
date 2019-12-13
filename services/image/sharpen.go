package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Sharpen(img *image.Image, opts options) (*image.NRGBA, error) {
	sigma, err := strconv.ParseFloat(opts.Get("sigma"), 64)
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("sigma must be a float"))
	}
	return imaging.Sharpen(*img, sigma), nil
}
