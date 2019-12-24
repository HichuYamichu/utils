package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
)

// Resize ImageService func
func Resize(img *image.Image, a *Args) (*image.NRGBA, error) {
	x, err := strconv.Atoi(a.X)
	if err != nil {
		return nil, fmt.Errorf("x position must be provided")
	}
	y, err := strconv.Atoi(a.Y)
	if err != nil {
		return nil, fmt.Errorf("y position must be provided")
	}
	filter, err := getFilterType(a.Filter)
	if err != nil {
		return nil, fmt.Errorf("invalid filter type")
	}

	return imaging.Resize(*img, x, y, *filter), nil
}
