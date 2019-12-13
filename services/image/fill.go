package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Fill(img *image.Image, opts options) (*image.NRGBA, error) {
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
		return nil, err
	}

	return imaging.Fill(*img, x, y, imaging.Center, *filter), nil
}
