package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snowballlanguage"
)

type SnowballTokenFilter struct {
	Language *snowballlanguage.SnowballLanguage `json:"language,omitempty"`
	Type     string                             `json:"type,omitempty"`
	Version  *string                            `json:"version,omitempty"`
}

func (s *SnowballTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SnowballTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSnowballTokenFilter() *SnowballTokenFilter { _ = "STUB: not implemented"; return nil }

type SnowballTokenFilterVariant interface {
	SnowballTokenFilterCaster() *SnowballTokenFilter
}

func (s *SnowballTokenFilter) SnowballTokenFilterCaster() *SnowballTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SnowballTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
