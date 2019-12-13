package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/hichuyamichu-me/utils/app/handler"
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
	a.srv.Handler = handler.New()
	a.srv.WriteTimeout = 15 * time.Second
	a.srv.ReadTimeout = 15 * time.Second
	return a
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
