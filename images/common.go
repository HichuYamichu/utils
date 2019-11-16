package images

import (
	"net/url"
	"strconv"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

type options struct {
	method string
	filter imaging.ResampleFilter
	x      int
	y      int
}

func parseOptions(form url.Values) (*options, *appErrors.Error) {
	method := form.Get("method")
	if method == "" {
		return nil, appErrors.New(400, "no resize method specyfied")
	}

	filter := form.Get("filter")
	if filter == "" {
		return nil, appErrors.New(400, "no resize filter specyfied")
	}

	x, err := strconv.Atoi(form.Get("x"))
	if err != nil {
		return nil, appErrors.New(400, "x must be an integer")
	}

	y, err := strconv.Atoi(form.Get("y"))
	if err != nil {
		return nil, appErrors.New(400, "x must be an integer")
	}

	var filterType imaging.ResampleFilter
	switch filter {
	case "Lanczos":
		filterType = imaging.Lanczos
	case "CatmullRom":
		filterType = imaging.CatmullRom
	case "MitchellNetravali":
		filterType = imaging.MitchellNetravali
	case "Linear":
		filterType = imaging.Linear
	case "Box":
		filterType = imaging.Box
	case "NearestNeighbor":
		filterType = imaging.NearestNeighbor
	default:
		return nil, appErrors.New(400, "illegal filter value")
	}

	return &options{method: method, filter: filterType, x: x, y: y}, nil
}
