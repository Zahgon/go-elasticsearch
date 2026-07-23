package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type XlmRobertaTokenizationConfig struct {
	DoLowerCase *bool `json:"do_lower_case,omitempty"`

	MaxSequenceLength *int `json:"max_sequence_length,omitempty"`

	Span *int `json:"span,omitempty"`

	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`

	WithSpecialTokens *bool `json:"with_special_tokens,omitempty"`
}

func (s *XlmRobertaTokenizationConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewXlmRobertaTokenizationConfig() *XlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

type XlmRobertaTokenizationConfigVariant interface {
	XlmRobertaTokenizationConfigCaster() *XlmRobertaTokenizationConfig
}

func (s *XlmRobertaTokenizationConfig) XlmRobertaTokenizationConfigCaster() *XlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
