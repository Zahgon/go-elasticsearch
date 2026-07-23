package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type _inferenceString struct {
	v *types.InferenceString
}

func NewInferenceString(type_ embeddingcontenttype.EmbeddingContentType, value string) *_inferenceString {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceString) Format(format embeddingcontentformat.EmbeddingContentFormat) *_inferenceString {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceString) Type(type_ embeddingcontenttype.EmbeddingContentType) *_inferenceString {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceString) Value(value string) *_inferenceString {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceString) InferenceStringCaster() *types.InferenceString {
	_ = "STUB: not implemented"
	return nil
}
