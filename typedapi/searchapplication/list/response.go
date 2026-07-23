package list

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count   int64                     `json:"count"`
	Results []types.SearchApplication `json:"results"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
