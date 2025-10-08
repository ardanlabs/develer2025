// Package web tiny little web framework
package web

import (
	"context"
	"net/http"
)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

type App struct {
	*http.ServeMux
}

func NewApp() *App {
	return &App{
		ServeMux: http.NewServeMux(),
	}
}

// HandleFunc IS MY API.
func (a *App) HandleFunc(pattern string, handler HandlerFunc) {
	h := func(w http.ResponseWriter, r *http.Request) {
		// I CAN DO ANYTHING HERE

		handler(r.Context(), w, r)

		// I CAN DO ANYTHING HERE
	}

	a.ServeMux.HandleFunc(pattern, h)
}
