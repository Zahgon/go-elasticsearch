package explain

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Explanation *types.ExplanationDetail `json:"explanation,omitempty"`
	Get         *types.InlineGet         `json:"get,omitempty"`
	Id_         string                   `json:"_id"`
	Index_      string                   `json:"_index"`
	Matched     bool                     `json:"matched"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
