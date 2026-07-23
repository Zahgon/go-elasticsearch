package resetfeatures

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Features []types.Feature `json:"features"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
