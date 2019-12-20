package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Fill(img *image.Image, a *Args) (*image.NRGBA, error) {
	x, err := strconv.Atoi(a.X)
	if err != nil {
		return nil, appErrors.Missing(fmt.Errorf("x position must be provided"))
	}
	y, err := strconv.Atoi(a.Y)
	if err != nil {
		return nil, appErrors.Missing(fmt.Errorf("y position must be provided"))
	}
	filter, err := getFilterType(a.Filter)
	if err != nil {
		return nil, err
	}

	return imaging.Fill(*img, x, y, imaging.Center, *filter), nil
}
