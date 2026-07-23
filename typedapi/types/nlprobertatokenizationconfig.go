package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type NlpRobertaTokenizationConfig struct {
	AddPrefixSpace *bool `json:"add_prefix_space,omitempty"`

	DoLowerCase *bool `json:"do_lower_case,omitempty"`

	MaxSequenceLength *int `json:"max_sequence_length,omitempty"`

	Span *int `json:"span,omitempty"`

	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`

	WithSpecialTokens *bool `json:"with_special_tokens,omitempty"`
}

func (s *NlpRobertaTokenizationConfig) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNlpRobertaTokenizationConfig() *NlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

type NlpRobertaTokenizationConfigVariant interface {
	NlpRobertaTokenizationConfigCaster() *NlpRobertaTokenizationConfig
}

func (s *NlpRobertaTokenizationConfig) NlpRobertaTokenizationConfigCaster() *NlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
