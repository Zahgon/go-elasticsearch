package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pinnedRetriever struct {
	v *types.PinnedRetriever
}

func NewPinnedRetriever(retriever types.RetrieverContainerVariant) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) Docs(docs ...types.SpecifiedDocumentVariant) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) DocsValues(docsvalues []types.SpecifiedDocument) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) Ids(ids ...string) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) RankWindowSize(rankwindowsize int) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) Retriever(retriever types.RetrieverContainerVariant) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) Filter(filters ...types.QueryVariant) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) MinScore(minscore float32) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) Name_(name_ string) *_pinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pinnedRetriever) PinnedRetrieverCaster() *types.PinnedRetriever {
	_ = "STUB: not implemented"
	return nil
}
