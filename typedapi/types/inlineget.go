package types

import (
	"encoding/json"
)

type InlineGet struct {
	Fields       map[string]json.RawMessage `json:"fields,omitempty"`
	Found        bool                       `json:"found"`
	Metadata     map[string]json.RawMessage `json:"-"`
	PrimaryTerm_ *int64                     `json:"_primary_term,omitempty"`
	Routing_     []string                   `json:"_routing,omitempty"`
	SeqNo_       *int64                     `json:"_seq_no,omitempty"`
	Source_      json.RawMessage            `json:"_source,omitempty"`
}

func (s *InlineGet) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s InlineGet) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewInlineGet() *InlineGet { _ = "STUB: not implemented"; return nil }
