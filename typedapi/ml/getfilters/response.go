package getfilters

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count   int64            `json:"count"`
	Filters []types.MLFilter `json:"filters"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
