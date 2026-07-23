package esapi

import (
	"context"
	"io"
	"net/http"
)

const (
	headerContentType = "Content-Type"
)

var (
	headerContentTypeJSON = []string{"application/json"}
)

type Request interface {
	Do(ctx context.Context, transport Transport) (*Response, error)
}

func newRequest(method, path string, body io.Reader) (*http.Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
