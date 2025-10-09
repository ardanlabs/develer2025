// Package hackapp yes yes.
package hackapp

import (
	"context"
	"math/rand/v2"
	"net/http"

	"github.com/ardanlabs/service/app/sdk/errs"
	"github.com/ardanlabs/service/foundation/web"
)

func Hack(ctx context.Context, w http.ResponseWriter, r *http.Request) web.Encoder {
	if n := rand.IntN(100); n%2 == 0 {
		return errs.Newf(errs.InvalidArgument, "THIS IS AN ERROR")
	}

	status := Status{
		Status: "ok",
	}

	return status
}
