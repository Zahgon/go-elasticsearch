package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsRegexp struct {
	v *types.IntervalsRegexp
}

func NewIntervalsRegexp(pattern string) *_intervalsRegexp { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRegexp) Analyzer(analyzer string) *_intervalsRegexp {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRegexp) Pattern(pattern string) *_intervalsRegexp {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRegexp) UseField(field string) *_intervalsRegexp {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRegexp) IntervalsCaster() *types.Intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRegexp) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRegexp) IntervalsRegexpCaster() *types.IntervalsRegexp {
	_ = "STUB: not implemented"
	return nil
}
