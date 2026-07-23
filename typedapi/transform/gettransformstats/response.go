package gettransformstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count      int64                  `json:"count"`
	Transforms []types.TransformStats `json:"transforms"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
