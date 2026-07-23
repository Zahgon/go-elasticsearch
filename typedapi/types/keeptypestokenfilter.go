package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/keeptypesmode"
)

type KeepTypesTokenFilter struct {
	Mode *keeptypesmode.KeepTypesMode `json:"mode,omitempty"`
	Type string                       `json:"type,omitempty"`

	Types   []string `json:"types"`
	Version *string  `json:"version,omitempty"`
}

func (s *KeepTypesTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s KeepTypesTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewKeepTypesTokenFilter() *KeepTypesTokenFilter { _ = "STUB: not implemented"; return nil }

type KeepTypesTokenFilterVariant interface {
	KeepTypesTokenFilterCaster() *KeepTypesTokenFilter
}

func (s *KeepTypesTokenFilter) KeepTypesTokenFilterCaster() *KeepTypesTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *KeepTypesTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
