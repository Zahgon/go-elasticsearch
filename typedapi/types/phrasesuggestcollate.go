package types

import (
	"encoding/json"
)

type PhraseSuggestCollate struct {
	Params map[string]json.RawMessage `json:"params,omitempty"`

	Prune *bool `json:"prune,omitempty"`

	Query PhraseSuggestCollateQuery `json:"query"`
}

func (s *PhraseSuggestCollate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPhraseSuggestCollate() *PhraseSuggestCollate { _ = "STUB: not implemented"; return nil }

type PhraseSuggestCollateVariant interface {
	PhraseSuggestCollateCaster() *PhraseSuggestCollate
}

func (s *PhraseSuggestCollate) PhraseSuggestCollateCaster() *PhraseSuggestCollate {
	_ = "STUB: not implemented"
	return nil
}
