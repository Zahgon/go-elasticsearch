package scriptspainlessexecute

import (
	"encoding/json"
)

type Response struct {
	Result json.RawMessage `json:"result,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
