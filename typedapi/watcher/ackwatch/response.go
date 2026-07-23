package ackwatch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Status types.WatchStatus `json:"status"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
