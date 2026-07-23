package mtermvectors

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Docs []types.TermVectorsResult `json:"docs"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
