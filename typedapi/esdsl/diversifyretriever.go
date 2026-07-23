package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/diversifyretrievertypes"
)

type _diversifyRetriever struct {
	v *types.DiversifyRetriever
}

func NewDiversifyRetriever(field string, retriever types.RetrieverContainerVariant, type_ diversifyretrievertypes.DiversifyRetrieverTypes) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Field(field string) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Lambda(lambda float32) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) QueryVector(queryvectors ...float32) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) QueryVectorBuilder(queryvectorbuilder types.QueryVectorBuilderVariant) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) RankWindowSize(rankwindowsize int) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Retriever(retriever types.RetrieverContainerVariant) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Size(size int) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Type(type_ diversifyretrievertypes.DiversifyRetrieverTypes) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Filter(filters ...types.QueryVariant) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) MinScore(minscore float32) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) Name_(name_ string) *_diversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_diversifyRetriever) DiversifyRetrieverCaster() *types.DiversifyRetriever {
	_ = "STUB: not implemented"
	return nil
}
