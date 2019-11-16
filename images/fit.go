package images

import (
	"image"
	"net/url"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Fit(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	opts, err := parseOptions(form)
	if err != nil {
		return nil, err
	}

	return imaging.Fit(*img, opts.x, opts.y, opts.filter), nil
}
