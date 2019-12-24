package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
)

// Gamma ImageService func
func Gamma(img *image.Image, a *Args) (*image.NRGBA, error) {
	gamma, err := strconv.ParseFloat(a.Value, 64)
	if err != nil {
		return nil, fmt.Errorf("gamma value must be a float")
	}
	return imaging.AdjustGamma(*img, gamma), nil
}
