package handlers

import (
	"bytes"
	"image"
	"image/jpeg"
	"net/http"
	"net/url"

	appErrors "github.com/hichuyamichu-me/utils/errors"
)

type ForImages func(img *image.Image, form url.Values) (*image.NRGBA, *appErrors.Error)

func (h ForImages) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, _, e := r.FormFile("file")
	if e != nil {
		http.Error(w, "no file provided", 400)
		return
	}

	img, _, e := image.Decode(file)
	if e != nil {
		http.Error(w, e.Error(), 400)
		return
	}

	res, err := h(&img, r.Form)
	if err != nil {
		http.Error(w, err.Msg, err.Code)
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
