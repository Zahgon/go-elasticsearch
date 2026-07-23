package putsettings

import (
	"encoding/json"
)

type Response struct {
	Acknowledged bool                       `json:"acknowledged"`
	Persistent   map[string]json.RawMessage `json:"persistent"`
	Transient    map[string]json.RawMessage `json:"transient"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
