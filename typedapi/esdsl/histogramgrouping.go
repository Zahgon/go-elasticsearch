package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _histogramGrouping struct {
	v *types.HistogramGrouping
}

func NewHistogramGrouping(interval int64) *_histogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramGrouping) Fields(fields ...string) *_histogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramGrouping) Interval(interval int64) *_histogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_histogramGrouping) HistogramGroupingCaster() *types.HistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}
