package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _standardRetriever struct {
	v *types.StandardRetriever
}

func NewStandardRetriever() *_standardRetriever { _ = "STUB: not implemented"; return nil }

func (s *_standardRetriever) Collapse(collapse types.FieldCollapseVariant) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) Query(query types.QueryVariant) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) SearchAfter(sortresults ...types.FieldValueVariant) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) SearchAfterValues(sortresultsvalues []types.FieldValue) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) Sort(sorts ...types.SortCombinationsVariant) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) SortValues(sortvalues []types.SortCombinations) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) TerminateAfter(terminateafter int) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) Filter(filters ...types.QueryVariant) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) MinScore(minscore float32) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) Name_(name_ string) *_standardRetriever {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) RetrieverContainerCaster() *types.RetrieverContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_standardRetriever) StandardRetrieverCaster() *types.StandardRetriever {
	_ = "STUB: not implemented"
	return nil
}
