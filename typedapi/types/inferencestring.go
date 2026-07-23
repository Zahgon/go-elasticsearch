package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type InferenceString struct {
	Format *embeddingcontentformat.EmbeddingContentFormat `json:"format,omitempty"`

	Type embeddingcontenttype.EmbeddingContentType `json:"type"`

	Value string `json:"value"`
}

func (s *InferenceString) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInferenceString() *InferenceString { _ = "STUB: not implemented"; return nil }

type InferenceStringVariant interface {
	InferenceStringCaster() *InferenceString
}

func (s *InferenceString) InferenceStringCaster() *InferenceString {
	_ = "STUB: not implemented"
	return nil
}

func (s *InferenceString) InferenceStringGroupCaster() *InferenceStringGroup {
	_ = "STUB: not implemented"
	return nil
}
