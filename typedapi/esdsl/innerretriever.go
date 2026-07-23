package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scorenormalizer"
)

type _innerRetriever struct {
	v *types.InnerRetriever
}

func NewInnerRetriever(normalizer scorenormalizer.ScoreNormalizer, retriever types.RetrieverContainerVariant, weight float32) *_innerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerRetriever) Normalizer(normalizer scorenormalizer.ScoreNormalizer) *_innerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerRetriever) Retriever(retriever types.RetrieverContainerVariant) *_innerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerRetriever) Weight(weight float32) *_innerRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_innerRetriever) InnerRetrieverCaster() *types.InnerRetriever {
	_ = "STUB: not implemented"
	return nil
}
