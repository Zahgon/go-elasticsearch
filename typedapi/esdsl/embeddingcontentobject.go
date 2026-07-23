package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _embeddingContentObject struct {
	v *types.EmbeddingContentObject
}

func NewEmbeddingContentObject() *_embeddingContentObject { _ = "STUB: not implemented"; return nil }

func (s *_embeddingContentObject) Content(embeddingcontentobjectgroups ...types.EmbeddingContentObjectItemVariant) *_embeddingContentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObject) ContentValues(embeddingcontentobjectgroupvalues []types.EmbeddingContentObjectItem) *_embeddingContentObject {
	_ = "STUB: not implemented"
	return nil
}

func (s *_embeddingContentObject) EmbeddingContentObjectCaster() *types.EmbeddingContentObject {
	_ = "STUB: not implemented"
	return nil
}
