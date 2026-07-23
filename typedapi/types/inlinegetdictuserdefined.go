package types

import (
	"encoding/json"
)

type InlineGetDictUserDefined struct {
	Fields                   map[string]json.RawMessage `json:"fields,omitempty"`
	Found                    bool                       `json:"found"`
	InlineGetDictUserDefined map[string]json.RawMessage `json:"-"`
	PrimaryTerm_             *int64                     `json:"_primary_term,omitempty"`
	Routing_                 []string                   `json:"_routing,omitempty"`
	SeqNo_                   *int64                     `json:"_seq_no,omitempty"`
	Source_                  map[string]json.RawMessage `json:"_source,omitempty"`
}

func (s *InlineGetDictUserDefined) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s InlineGetDictUserDefined) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewInlineGetDictUserDefined() *InlineGetDictUserDefined { _ = "STUB: not implemented"; return nil }
