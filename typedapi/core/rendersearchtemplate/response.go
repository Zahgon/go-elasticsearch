package rendersearchtemplate

import (
	"encoding/json"
)

type Response struct {
	TemplateOutput map[string]json.RawMessage `json:"template_output"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
