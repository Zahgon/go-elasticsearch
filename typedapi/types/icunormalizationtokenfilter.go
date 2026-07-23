package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type IcuNormalizationTokenFilter struct {
	Name    icunormalizationtype.IcuNormalizationType `json:"name"`
	Type    string                                    `json:"type,omitempty"`
	Version *string                                   `json:"version,omitempty"`
}

func (s *IcuNormalizationTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuNormalizationTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuNormalizationTokenFilter() *IcuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type IcuNormalizationTokenFilterVariant interface {
	IcuNormalizationTokenFilterCaster() *IcuNormalizationTokenFilter
}

func (s *IcuNormalizationTokenFilter) IcuNormalizationTokenFilterCaster() *IcuNormalizationTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuNormalizationTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
