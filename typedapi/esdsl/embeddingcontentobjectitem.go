package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type _embeddingContentObjectItem struct {
	v *types.EmbeddingContentObjectItem
}

func NewEmbeddingContentObjectItem(type_ embeddingcontenttype.EmbeddingContentType, value string) *_embeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObjectItem) Format(format embeddingcontentformat.EmbeddingContentFormat) *_embeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObjectItem) Type(type_ embeddingcontenttype.EmbeddingContentType) *_embeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObjectItem) Value(value string) *_embeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObjectItem) EmbeddingContentObjectItemCaster() *types.EmbeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}
