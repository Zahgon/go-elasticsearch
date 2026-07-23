package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnSearch struct {
	v *types.KnnSearch
}

func NewKnnSearch() *_knnSearch { _ = "STUB: not implemented"; return nil }

func (s *_knnSearch) Boost(boost float32) *_knnSearch { _ = "STUB: not implemented"; return nil }

func (s *_knnSearch) Field(field string) *_knnSearch { _ = "STUB: not implemented"; return nil }

func (s *_knnSearch) Filter(filters ...types.QueryVariant) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) InnerHits(innerhits types.InnerHitsVariant) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) K(k int) *_knnSearch { _ = "STUB: not implemented"; return nil }

func (s *_knnSearch) NumCandidates(numcandidates int) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) QueryName_(queryname_ string) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) QueryVector(queryvectors ...float32) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) Similarity(similarity float32) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) VisitPercentage(visitpercentage float32) *_knnSearch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnSearch) KnnSearchCaster() *types.KnnSearch { _ = "STUB: not implemented"; return nil }
