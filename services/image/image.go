package image

import (
	"fmt"
	"net/url"

	"github.com/disintegration/imaging"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

type options = url.Values

func getFilterType(typeName string) (*imaging.ResampleFilter, error) {
	var filterType imaging.ResampleFilter
	switch typeName {
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
		return nil, appErrors.InvalidType(fmt.Errorf("invalid filter type"))
	}
	return &filterType, nil
}
