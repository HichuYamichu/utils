package handler

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
	"net/url"

	"github.com/hichuyamichu-me/utils/app/store"
	appErrors "github.com/hichuyamichu-me/utils/errors"
)

type imageHandler struct {
	store *store.Store
	sf    imageServiceFunc
}

type imageServiceFunc func(img *image.Image, opts url.Values) (*image.NRGBA, error)

func newImageHandler(store *store.Store, sf imageServiceFunc) *imageHandler {
	return &imageHandler{store: store, sf: sf}
}

func (h imageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "no file provided", 400)
		return
	}

	img, _, e := image.Decode(file)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}

	res, err := h.sf(&img, r.Form)
	switch err.(type) {
	case appErrors.Missing:
	case appErrors.InvalidType:
		http.Error(w, err.Error(), 400)
		return
	case appErrors.NotFound:
		http.Error(w, err.Error(), 404)
		return
	}

	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, res, nil); err != nil {
		http.Error(w, "unable to encode image", 500)
		return
	}

	if _, err := w.Write(buffer.Bytes()); err != nil {
		http.Error(w, "unable to write image", 500)
	}
}
