package esapi

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Header     http.Header

	Body io.ReadCloser
}

func (r *Response) String() string { _ = "STUB: not implemented"; return "" }

func (r *Response) Status() string { _ = "STUB: not implemented"; return "" }

func (r *Response) IsError() bool { _ = "STUB: not implemented"; return false }

func (r *Response) Warnings() []string { _ = "STUB: not implemented"; return nil }

func (r *Response) HasWarnings() bool { _ = "STUB: not implemented"; return false }
