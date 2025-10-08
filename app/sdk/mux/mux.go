// Package mux is yada yada.
package mux

import (
	"github.com/ardanlabs/service/app/domain/hackapp"
	"github.com/ardanlabs/service/app/sdk/mid"
	"github.com/ardanlabs/service/foundation/logger"
	"github.com/ardanlabs/service/foundation/web"
)

func WebAPI(log *logger.Logger) *web.App {
	mux := web.NewApp(log.Info, mid.Logger(log))

	mux.HandleFunc("GET /test", hackapp.Hack)

	return mux
}
