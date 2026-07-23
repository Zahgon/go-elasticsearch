package renderquery

import (
	"encoding/json"
)

type Request struct {
	Params map[string]json.RawMessage `json:"params,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
