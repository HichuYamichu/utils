package image

import (
	"fmt"
	"image"
	"net/url"

	"github.com/disintegration/imaging"
)

// Args arguments for ImageService functions
type Args struct {
	img    *image.Image
	X      string
	Y      string
	Filter string
	Value  string
}

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
		return nil, fmt.Errorf("invalid filter type")
	}
	return &filterType, nil
}
