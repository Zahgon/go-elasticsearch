package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type IcuNormalizationCharFilter struct {
	Mode             *icunormalizationmode.IcuNormalizationMode `json:"mode,omitempty"`
	Name             *icunormalizationtype.IcuNormalizationType `json:"name,omitempty"`
	Type             string                                     `json:"type,omitempty"`
	UnicodeSetFilter *string                                    `json:"unicode_set_filter,omitempty"`
	Version          *string                                    `json:"version,omitempty"`
}

func (s *IcuNormalizationCharFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuNormalizationCharFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuNormalizationCharFilter() *IcuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

type IcuNormalizationCharFilterVariant interface {
	IcuNormalizationCharFilterCaster() *IcuNormalizationCharFilter
}

func (s *IcuNormalizationCharFilter) IcuNormalizationCharFilterCaster() *IcuNormalizationCharFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuNormalizationCharFilter) CharFilterDefinitionCaster() *CharFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
