package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rescorerRetriever struct {
	v *types.RescorerRetriever
}

func NewRescorerRetriever(retriever types.RetrieverContainerVariant) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) Rescore(rescores ...types.RescoreVariant) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) Retriever(retriever types.RetrieverContainerVariant) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) Filter(filters ...types.QueryVariant) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) MinScore(minscore float32) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) Name_(name_ string) *_rescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rescorerRetriever) RescorerRetrieverCaster() *types.RescorerRetriever {
	_ = "STUB: not implemented"
	return nil
}
