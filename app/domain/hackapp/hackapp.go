// Package hackapp yes yes.
package hackapp

import (
	"context"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	status := Status{
		Status: "ok",
	}

	return status
}
