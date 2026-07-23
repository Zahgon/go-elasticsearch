package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type NlpTokenizationUpdateOptions struct {
	Span *int `json:"span,omitempty"`

	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`
}

func (s *NlpTokenizationUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNlpTokenizationUpdateOptions() *NlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type NlpTokenizationUpdateOptionsVariant interface {
	NlpTokenizationUpdateOptionsCaster() *NlpTokenizationUpdateOptions
}

func (s *NlpTokenizationUpdateOptions) NlpTokenizationUpdateOptionsCaster() *NlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
