// Package mux is yada yada.
package mux

import (
	"github.com/ardanlabs/service/app/domain/hackapp"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI() *web.App {
	mux := web.NewApp()

	mux.HandleFunc("GET /test", hackapp.Hack)

	return mux
}
