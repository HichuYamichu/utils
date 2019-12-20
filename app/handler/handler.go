package handler

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hichuyamichu-me/utils/app/store"
	"github.com/hichuyamichu-me/utils/services/image"
)

func New(s *store.Store) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/images", func(r chi.Router) {
			r.Method("POST", "/fit", newImageHandler(s, image.Fit))
			r.Method("POST", "/fill", newImageHandler(s, image.Fill))
			r.Method("POST", "/resize", newImageHandler(s, image.Resize))
			r.Method("POST", "/blurr", newImageHandler(s, image.Blurr))
			r.Method("POST", "/saturate", newImageHandler(s, image.Saturation))
			r.Method("POST", "/sharpen", newImageHandler(s, image.Sharpen))
			r.Method("POST", "/gamma", newImageHandler(s, image.Gamma))
			r.Method("POST", "/contrast", newImageHandler(s, image.Contrast))
		})
	})

	return r
}

type handler struct {
	store *store.Store
}
