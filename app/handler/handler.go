package handler

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hichuyamichu-me/utils/services/image"
)

func New() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/images", func(r chi.Router) {
			r.Method("POST", "/fit", imageHandler(image.Fit))
			r.Method("POST", "/fill", imageHandler(image.Fill))
			r.Method("POST", "/resize", imageHandler(image.Resize))
			r.Method("POST", "/blurr", imageHandler(image.Blurr))
			r.Method("POST", "/saturate", imageHandler(image.Saturation))
			r.Method("POST", "/sharpen", imageHandler(image.Sharpen))
			r.Method("POST", "/gamma", imageHandler(image.Gamma))
			r.Method("POST", "/contrast", imageHandler(image.Contrast))
		})
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "./client/UtilsClient/dist")
	FileServer(r, "/", http.Dir(filesDir))
	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))
}
