package putsettings

import (
	"encoding/json"
)

type Request struct {
	Persistent map[string]json.RawMessage `json:"persistent,omitempty"`

	Transient map[string]json.RawMessage `json:"transient,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
