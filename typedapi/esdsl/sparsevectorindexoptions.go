package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sparseVectorIndexOptions struct {
	v *types.SparseVectorIndexOptions
}

func NewSparseVectorIndexOptions() *_sparseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorIndexOptions) Prune(prune bool) *_sparseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorIndexOptions) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_sparseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorIndexOptions) SparseVectorIndexOptionsCaster() *types.SparseVectorIndexOptions {
	_ = "STUB: not implemented"
	return nil
}
