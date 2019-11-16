package images

import (
	"fmt"
	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
	"image"
	"net/url"
	"strconv"
)

func Blurr(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	sigma, err := strconv.ParseFloat(form.Get("sigma"), 64)
	fmt.Println(sigma)
	if err != nil {
		return nil, appErrors.New(400, "sigma must be a float")
	}
	return imaging.Blur(*img, sigma), nil
}
