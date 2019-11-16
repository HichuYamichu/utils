package images

import (
	"image"
	"net/url"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

func Fill(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error) {
	opts, err := parseOptions(form)
	if err != nil {
		return nil, err
	}

	return imaging.Fill(*img, opts.x, opts.y, imaging.Center, opts.filter), nil
}
