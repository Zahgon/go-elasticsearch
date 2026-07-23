package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type _linearRetriever struct {
	v *types.LinearRetriever
}

func NewLinearRetriever() *_linearRetriever { _ = "STUB: not implemented"; return nil }

func (s *_linearRetriever) Fields(fields ...string) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) Normalizer(normalizer scorenormalizer.ScoreNormalizer) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) Query(query string) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) RankWindowSize(rankwindowsize int) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) Retrievers(retrievers ...types.InnerRetrieverVariant) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) RetrieversValues(retrieversvalues []types.InnerRetriever) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) Filter(filters ...types.QueryVariant) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) MinScore(minscore float32) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) Name_(name_ string) *_linearRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearRetriever) LinearRetrieverCaster() *types.LinearRetriever {
	_ = "STUB: not implemented"
	return nil
}
