package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Blurr(img *image.Image, opts options) (*image.NRGBA, error) {
	sigma, err := strconv.ParseFloat(opts.Get("sigma"), 64)
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("sigma must be a float"))
	}
	return imaging.Blur(*img, sigma), nil
}
