package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
)

// Contrast ImageService func
func Contrast(img *image.Image, a *Args) (*image.NRGBA, error) {
	contrast, err := strconv.ParseFloat(a.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("contrast value must be a float")
	}
	return imaging.AdjustContrast(*img, contrast), nil
}
