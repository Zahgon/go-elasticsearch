package stats

import (
	"encoding/json"
)

type Response struct {
	Stats json.RawMessage `json:"stats,omitempty"`
	Total json.RawMessage `json:"total,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
