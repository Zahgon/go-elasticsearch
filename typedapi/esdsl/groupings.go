package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _groupings struct {
	v *types.Groupings
}

func NewGroupings() *_groupings { _ = "STUB: not implemented"; return nil }

func (s *_groupings) DateHistogram(datehistogram types.DateHistogramGroupingVariant) *_groupings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groupings) Histogram(histogram types.HistogramGroupingVariant) *_groupings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groupings) Terms(terms types.TermsGroupingVariant) *_groupings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_groupings) GroupingsCaster() *types.Groupings { _ = "STUB: not implemented"; return nil }
