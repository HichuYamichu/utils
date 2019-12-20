package image

import (
	"fmt"
	"image"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Saturation(img *image.Image, a *Args) (*image.NRGBA, error) {
	saturation, err := strconv.ParseFloat(a.Value, 64)
	if err != nil {
		return nil, appErrors.InvalidType(fmt.Errorf("saturation value must be a float"))
	}
	return imaging.AdjustSaturation(*img, saturation), nil
}
