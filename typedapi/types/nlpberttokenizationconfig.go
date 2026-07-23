package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type NlpBertTokenizationConfig struct {
	DoLowerCase *bool `json:"do_lower_case,omitempty"`

	MaxSequenceLength *int `json:"max_sequence_length,omitempty"`

	Span *int `json:"span,omitempty"`

	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`

	WithSpecialTokens *bool `json:"with_special_tokens,omitempty"`
}

func (s *NlpBertTokenizationConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNlpBertTokenizationConfig() *NlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

type NlpBertTokenizationConfigVariant interface {
	NlpBertTokenizationConfigCaster() *NlpBertTokenizationConfig
}

func (s *NlpBertTokenizationConfig) NlpBertTokenizationConfigCaster() *NlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
