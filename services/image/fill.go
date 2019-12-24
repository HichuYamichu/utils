package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
)

// Fill ImageService func
func Fill(img *image.Image, a *Args) (*image.NRGBA, error) {
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
		return nil, err
	}

	return imaging.Fill(*img, x, y, imaging.Center, *filter), nil
}
