package getsettings

import (
	"encoding/json"
)

type Response struct {
	Defaults map[string]json.RawMessage `json:"defaults,omitempty"`

	Persistent map[string]json.RawMessage `json:"persistent"`

	Transient map[string]json.RawMessage `json:"transient"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
