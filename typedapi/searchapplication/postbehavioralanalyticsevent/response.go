package postbehavioralanalyticsevent

import (
	"encoding/json"
)

type Response struct {
	Accepted bool            `json:"accepted"`
	Event    json.RawMessage `json:"event,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
