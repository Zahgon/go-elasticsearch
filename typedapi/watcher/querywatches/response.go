package querywatches

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	Watches []types.QueryWatch `json:"watches"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
