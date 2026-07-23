package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnSearchQuery struct {
	v *types.KnnSearchQuery
}

func NewKnnSearchQuery(k int, numcandidates int) *_knnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearchQuery) Field(field string) *_knnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearchQuery) K(k int) *_knnSearchQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnSearchQuery) NumCandidates(numcandidates int) *_knnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearchQuery) QueryVector(queryvectors ...float32) *_knnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearchQuery) KnnSearchQueryCaster() *types.KnnSearchQuery {
	_ = "STUB: not implemented"
	return nil
}
