package typedapi

import (
	"context"
	"net/http"
)

type Request interface {
	Perform(ctx context.Context) (*http.Response, error)
}
