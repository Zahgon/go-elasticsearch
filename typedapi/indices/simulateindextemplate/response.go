package simulateindextemplate

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Overlapping []types.Overlapping `json:"overlapping,omitempty"`
	Template    types.Template      `json:"template"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
