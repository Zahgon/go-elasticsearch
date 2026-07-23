package queryuser

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	Total int `json:"total"`

	Users []types.QueryUser `json:"users"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
