package get

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Endpoints []types.InferenceEndpointInfo `json:"endpoints"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
