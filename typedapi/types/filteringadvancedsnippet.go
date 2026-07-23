package types

import (
	"encoding/json"
)

type FilteringAdvancedSnippet struct {
	CreatedAt DateTime        `json:"created_at,omitempty"`
	UpdatedAt DateTime        `json:"updated_at,omitempty"`
	Value     json.RawMessage `json:"value,omitempty"`
}

func (s *FilteringAdvancedSnippet) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFilteringAdvancedSnippet() *FilteringAdvancedSnippet { _ = "STUB: not implemented"; return nil }

type FilteringAdvancedSnippetVariant interface {
	FilteringAdvancedSnippetCaster() *FilteringAdvancedSnippet
}

func (s *FilteringAdvancedSnippet) FilteringAdvancedSnippetCaster() *FilteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}
