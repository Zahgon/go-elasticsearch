package types

import (
	"encoding/json"
)

type GetResult struct {
	Fields map[string]json.RawMessage `json:"fields,omitempty"`

	Found bool `json:"found"`

	Id_      string   `json:"_id"`
	Ignored_ []string `json:"_ignored,omitempty"`

	Index_ string `json:"_index"`

	PrimaryTerm_ *int64 `json:"_primary_term,omitempty"`

	Routing_ *string `json:"_routing,omitempty"`

	SeqNo_ *int64 `json:"_seq_no,omitempty"`

	Source_ json.RawMessage `json:"_source,omitempty"`

	Version_ *int64 `json:"_version,omitempty"`
}

func (s *GetResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGetResult() *GetResult { _ = "STUB: not implemented"; return nil }
