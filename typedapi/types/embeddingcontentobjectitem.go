package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontentformat"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/embeddingcontenttype"
)

type EmbeddingContentObjectItem struct {
	Format *embeddingcontentformat.EmbeddingContentFormat `json:"format,omitempty"`

	Type embeddingcontenttype.EmbeddingContentType `json:"type"`

	Value string `json:"value"`
}

func (s *EmbeddingContentObjectItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewEmbeddingContentObjectItem() *EmbeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

type EmbeddingContentObjectItemVariant interface {
	EmbeddingContentObjectItemCaster() *EmbeddingContentObjectItem
}

func (s *EmbeddingContentObjectItem) EmbeddingContentObjectItemCaster() *EmbeddingContentObjectItem {
	_ = "STUB: not implemented"
	return nil
}

func (s *EmbeddingContentObjectItem) EmbeddingContentObjectGroupCaster() *EmbeddingContentObjectGroup {
	_ = "STUB: not implemented"
	return nil
}
