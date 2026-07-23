package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateHistogramGrouping struct {
	v *types.DateHistogramGrouping
}

func NewDateHistogramGrouping() *_dateHistogramGrouping { _ = "STUB: not implemented"; return nil }

func (s *_dateHistogramGrouping) CalendarInterval(duration types.DurationVariant) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) Delay(duration types.DurationVariant) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) Field(field string) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) FixedInterval(duration types.DurationVariant) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) Format(format string) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) Interval(duration types.DurationVariant) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) TimeZone(timezone string) *_dateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramGrouping) DateHistogramGroupingCaster() *types.DateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}
