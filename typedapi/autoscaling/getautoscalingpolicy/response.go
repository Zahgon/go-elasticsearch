package getautoscalingpolicy

import (
	"encoding/json"
)

type Response struct {
	Deciders map[string]json.RawMessage `json:"deciders"`
	Roles    []string                   `json:"roles"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
