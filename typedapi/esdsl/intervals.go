package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervals struct {
	v *types.Intervals
}

func NewIntervals() *_intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervals) AllOf(allof types.IntervalsAllOfVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) AnyOf(anyof types.IntervalsAnyOfVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Fuzzy(fuzzy types.IntervalsFuzzyVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Match(match types.IntervalsMatchVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Prefix(prefix types.IntervalsPrefixVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Range(range_ types.IntervalsRangeVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Regexp(regexp types.IntervalsRegexpVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) Wildcard(wildcard types.IntervalsWildcardVariant) *_intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervals) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }
