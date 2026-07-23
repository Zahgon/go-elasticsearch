package queryrole

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	Roles []types.QueryRole `json:"roles"`

	Total int `json:"total"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
