package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rRFRetriever struct {
	v *types.RRFRetriever
}

func NewRRFRetriever() *_rRFRetriever { _ = "STUB: not implemented"; return nil }

func (s *_rRFRetriever) Fields(fields ...string) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) Query(query string) *_rRFRetriever { _ = "STUB: not implemented"; return nil }

func (s *_rRFRetriever) RankConstant(rankconstant int) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) RankWindowSize(rankwindowsize int) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) Retrievers(retrievers ...types.RRFRetrieverEntryVariant) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) RetrieversValues(retrieversvalues []types.RRFRetrieverEntry) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) Filter(filters ...types.QueryVariant) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) MinScore(minscore float32) *_rRFRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) Name_(name_ string) *_rRFRetriever { _ = "STUB: not implemented"; return nil }

func (s *_rRFRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rRFRetriever) RRFRetrieverCaster() *types.RRFRetriever {
	_ = "STUB: not implemented"
	return nil
}
