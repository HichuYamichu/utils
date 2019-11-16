package images

import (
	"image"
	"net/url"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Resize(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	opts, err := parseOptions(form)
	if err != nil {
		return nil, err
	}

	return imaging.Resize(*img, opts.x, opts.y, opts.filter), nil
}
