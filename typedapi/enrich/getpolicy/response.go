package getpolicy

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Policies []types.Summary `json:"policies"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
