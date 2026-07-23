package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _semanticTextIndexOptions struct {
	v *types.SemanticTextIndexOptions
}

func NewSemanticTextIndexOptions() *_semanticTextIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextIndexOptions) DenseVector(densevector types.DenseVectorIndexOptionsVariant) *_semanticTextIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextIndexOptions) SparseVector(sparsevector types.SparseVectorIndexOptionsVariant) *_semanticTextIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_semanticTextIndexOptions) SemanticTextIndexOptionsCaster() *types.SemanticTextIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
