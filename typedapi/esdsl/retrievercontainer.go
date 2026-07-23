package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _retrieverContainer struct {
	v *types.RetrieverContainer
}

func NewRetrieverContainer() *_retrieverContainer { _ = "STUB: not implemented"; return nil }

func (s *_retrieverContainer) Diversify(diversify types.DiversifyRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Knn(knn types.KnnRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Linear(linear types.LinearRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Pinned(pinned types.PinnedRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Rescorer(rescorer types.RescorerRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Rrf(rrf types.RRFRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Rule(rule types.RuleRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) Standard(standard types.StandardRetrieverVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) TextSimilarityReranker(textsimilarityreranker types.TextSimilarityRerankerVariant) *_retrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retrieverContainer) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}
