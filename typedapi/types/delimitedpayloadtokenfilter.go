package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/delimitedpayloadencoding"
)

type DelimitedPayloadTokenFilter struct {
	Delimiter *string `json:"delimiter,omitempty"`

	Encoding *delimitedpayloadencoding.DelimitedPayloadEncoding `json:"encoding,omitempty"`
	Type     string                                             `json:"type,omitempty"`
	Version  *string                                            `json:"version,omitempty"`
}

func (s *DelimitedPayloadTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DelimitedPayloadTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDelimitedPayloadTokenFilter() *DelimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

type DelimitedPayloadTokenFilterVariant interface {
	DelimitedPayloadTokenFilterCaster() *DelimitedPayloadTokenFilter
}

func (s *DelimitedPayloadTokenFilter) DelimitedPayloadTokenFilterCaster() *DelimitedPayloadTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *DelimitedPayloadTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
