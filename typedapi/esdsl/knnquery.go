package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnQuery struct {
	v *types.KnnQuery
}

func NewKnnQuery() *_knnQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) Field(field string) *_knnQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) Filter(filters ...types.QueryVariant) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) K(k int) *_knnQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) NumCandidates(numcandidates int) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) QueryVector(queryvectors ...float32) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) Similarity(similarity float32) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) VisitPercentage(visitpercentage float32) *_knnQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnQuery) Boost(boost float32) *_knnQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) QueryName_(queryname_ string) *_knnQuery { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func (s *_knnQuery) KnnQueryCaster() *types.KnnQuery { _ = "STUB: not implemented"; return nil }
