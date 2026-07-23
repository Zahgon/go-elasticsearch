package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _filteringAdvancedSnippet struct {
	v *types.FilteringAdvancedSnippet
}

func NewFilteringAdvancedSnippet(value json.RawMessage) *_filteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringAdvancedSnippet) CreatedAt(datetime types.DateTimeVariant) *_filteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringAdvancedSnippet) UpdatedAt(datetime types.DateTimeVariant) *_filteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringAdvancedSnippet) Value(value json.RawMessage) *_filteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}

func (s *_filteringAdvancedSnippet) FilteringAdvancedSnippetCaster() *types.FilteringAdvancedSnippet {
	_ = "STUB: not implemented"
	return nil
}
