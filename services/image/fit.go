package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Fit(img *image.Image, opts options) (*image.NRGBA, error) {
	x, err := strconv.Atoi(opts.Get("x"))
	if err != nil {
		return nil, appErrors.Missing(fmt.Errorf("x position must be provided"))
	}
	y, err := strconv.Atoi(opts.Get("y"))
	if err != nil {
		return nil, appErrors.Missing(fmt.Errorf("y position must be provided"))
	}
	filter, err := getFilterType(opts.Get("filter"))
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("invalid filter type"))
	}

	return imaging.Fit(*img, x, y, *filter), nil
}
