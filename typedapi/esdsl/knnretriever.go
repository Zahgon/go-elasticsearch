package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _knnRetriever struct {
	v *types.KnnRetriever
}

func NewKnnRetriever(field string, k int, numcandidates int) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) Field(field string) *_knnRetriever { _ = "STUB: not implemented"; return nil }

func (s *_knnRetriever) K(k int) *_knnRetriever { _ = "STUB: not implemented"; return nil }

func (s *_knnRetriever) NumCandidates(numcandidates int) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) QueryVector(queryvectors ...float32) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) RescoreVector(rescorevector types.RescoreVectorVariant) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) Similarity(similarity float32) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) VisitPercentage(visitpercentage float32) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) Filter(filters ...types.QueryVariant) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) MinScore(minscore float32) *_knnRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) Name_(name_ string) *_knnRetriever { _ = "STUB: not implemented"; return nil }

func (s *_knnRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_knnRetriever) KnnRetrieverCaster() *types.KnnRetriever {
	_ = "STUB: not implemented"
	return nil
}
