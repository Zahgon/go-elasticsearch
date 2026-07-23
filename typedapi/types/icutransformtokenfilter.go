package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icutransformdirection"
)

type IcuTransformTokenFilter struct {
	Dir     *icutransformdirection.IcuTransformDirection `json:"dir,omitempty"`
	Id      string                                       `json:"id"`
	Type    string                                       `json:"type,omitempty"`
	Version *string                                      `json:"version,omitempty"`
}

func (s *IcuTransformTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s IcuTransformTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewIcuTransformTokenFilter() *IcuTransformTokenFilter { _ = "STUB: not implemented"; return nil }

type IcuTransformTokenFilterVariant interface {
	IcuTransformTokenFilterCaster() *IcuTransformTokenFilter
}

func (s *IcuTransformTokenFilter) IcuTransformTokenFilterCaster() *IcuTransformTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *IcuTransformTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
