package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sparseVectorQuery struct {
	v *types.SparseVectorQuery
}

func NewSparseVectorQuery() *_sparseVectorQuery { _ = "STUB: not implemented"; return nil }

func (s *_sparseVectorQuery) Field(field string) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) InferenceId(id string) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) Prune(prune bool) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) PruningConfig(pruningconfig types.TokenPruningConfigVariant) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) Query(query string) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) QueryVector(queryvector map[string]float32) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) AddQueryVector(key string, value float32) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) Boost(boost float32) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) QueryName_(queryname_ string) *_sparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sparseVectorQuery) SparseVectorQueryCaster() *types.SparseVectorQuery {
	_ = "STUB: not implemented"
	return nil
}
