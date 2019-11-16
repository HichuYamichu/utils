package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/hichuyamichu-me/utils/handlers"
	"github.com/hichuyamichu-me/utils/images"
)

// App : Application struct
type App struct {
	srv *http.Server
}

// New : Initialize new server instance
func New(host, port string) *App {
	a := &App{}
	a.srv = &http.Server{}
	a.srv.Addr = fmt.Sprintf("%s:%s", host, port)
	a.srv.Handler = a.setupHandler()
	a.srv.WriteTimeout = 15 * time.Second
	a.srv.ReadTimeout = 15 * time.Second
	return a
}

func (a *App) setupHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Route("/images", func(r chi.Router) {
			r.Method("POST", "/fit", handlers.ForImages(images.Fit))
			r.Method("POST", "/fill", handlers.ForImages(images.Fill))
			r.Method("POST", "/resize", handlers.ForImages(images.Resize))
			r.Method("POST", "/blurr", handlers.ForImages(images.Blurr))
			r.Method("POST", "/saturate", handlers.ForImages(images.Saturation))
			r.Method("POST", "/sharpen", handlers.ForImages(images.Sharpen))
			r.Method("POST", "/gamma", handlers.ForImages(images.Gamma))
			r.Method("POST", "/contrast", handlers.ForImages(images.Contrast))
		})
	})

	fs := http.FileServer(http.Dir("static"))
	r.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat("static" + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
	return r
}

// Run : Starts the app
func (a *App) Run() {
	log.Printf("Listening on: http://%s\n", a.srv.Addr)
	log.Fatal(a.srv.ListenAndServe())
}

// Shutdown : Stops the app
func (a *App) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := a.srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
}
