package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Contrast(img *image.Image, opts options) (*image.NRGBA, error) {
	contrast, err := strconv.ParseFloat(opts.Get("contrast"), 64)
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("contrast value must be a float"))
	}
	return imaging.AdjustContrast(*img, contrast), nil
}
