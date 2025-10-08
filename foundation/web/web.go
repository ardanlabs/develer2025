// Package web tiny little web framework
package web

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Encoder interface {
	Encode() (data []byte, contentType string, err error)
}

type Logger func(ctx context.Context, msg string, args ...any)

type HandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request) Encoder

type App struct {
	log Logger
	*http.ServeMux
	mw []MidFunc
}

func NewApp(log Logger, mw ...MidFunc) *App {
	return &App{
		log:      log,
		ServeMux: http.NewServeMux(),
		mw:       mw,
	}
}

// HandleFunc IS MY API.
func (a *App) HandleFunc(pattern string, handlerFunc HandlerFunc, mw ...MidFunc) {
	handlerFunc = wrapMiddleware(mw, handlerFunc)
	handlerFunc = wrapMiddleware(a.mw, handlerFunc)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := setTraceID(r.Context(), uuid.New())

		resp := handlerFunc(ctx, w, r)

		if err := Respond(ctx, w, resp); err != nil {
			a.log(ctx, "web-respond", "ERROR", err)
			return
		}
	}

	a.ServeMux.HandleFunc(pattern, h)
}
