package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/lowercasetokenfilterlanguages"
)

type LowercaseTokenFilter struct {
	Language *lowercasetokenfilterlanguages.LowercaseTokenFilterLanguages `json:"language,omitempty"`
	Type     string                                                       `json:"type,omitempty"`
	Version  *string                                                      `json:"version,omitempty"`
}

func (s *LowercaseTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LowercaseTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLowercaseTokenFilter() *LowercaseTokenFilter { _ = "STUB: not implemented"; return nil }

type LowercaseTokenFilterVariant interface {
	LowercaseTokenFilterCaster() *LowercaseTokenFilter
}

func (s *LowercaseTokenFilter) LowercaseTokenFilterCaster() *LowercaseTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *LowercaseTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
