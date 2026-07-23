package getnodestats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Total                  types.TransformNodeStats            `json:"total"`
	TransformNodeFullStats map[string]types.TransformNodeStats `json:"-"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
