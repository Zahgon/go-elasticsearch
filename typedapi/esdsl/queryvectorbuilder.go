package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _queryVectorBuilder struct {
	v *types.QueryVectorBuilder
}

func NewQueryVectorBuilder() *_queryVectorBuilder { _ = "STUB: not implemented"; return nil }

func (s *_queryVectorBuilder) Embedding(embedding types.EmbeddingVariant) *_queryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryVectorBuilder) Lookup(lookup types.LookupQueryVectorBuilderVariant) *_queryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryVectorBuilder) TextEmbedding(textembedding types.TextEmbeddingVariant) *_queryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}

func (s *_queryVectorBuilder) QueryVectorBuilderCaster() *types.QueryVectorBuilder {
	_ = "STUB: not implemented"
	return nil
}
