package mount

import (
	"encoding/json"
)

type Request struct {
	IgnoreIndexSettings []string `json:"ignore_index_settings,omitempty"`

	Index string `json:"index"`

	IndexSettings map[string]json.RawMessage `json:"index_settings,omitempty"`

	RenamedIndex *string `json:"renamed_index,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
