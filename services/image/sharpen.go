package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
)

// Sharpen ImageService func
func Sharpen(img *image.Image, a *Args) (*image.NRGBA, error) {
	sigma, err := strconv.ParseFloat(a.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("sigma must be a float")
	}
	return imaging.Sharpen(*img, sigma), nil
}
