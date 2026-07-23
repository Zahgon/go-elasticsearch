package updateuserprofiledata

import (
	"encoding/json"
)

type Request struct {
	Data map[string]json.RawMessage `json:"data,omitempty"`

	Labels map[string]json.RawMessage `json:"labels,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
