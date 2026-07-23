package analyze

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Attributes []string `json:"attributes,omitempty"`

	CharFilter []types.CharFilter `json:"char_filter,omitempty"`

	Explain *bool `json:"explain,omitempty"`

	Field *string `json:"field,omitempty"`

	Filter []types.TokenFilter `json:"filter,omitempty"`

	Normalizer *string `json:"normalizer,omitempty"`

	Text []string `json:"text,omitempty"`

	Tokenizer types.Tokenizer `json:"tokenizer,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
