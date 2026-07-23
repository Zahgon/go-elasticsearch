package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsAnyOf struct {
	v *types.IntervalsAnyOf
}

func NewIntervalsAnyOf() *_intervalsAnyOf { _ = "STUB: not implemented"; return nil }

func (s *_intervalsAnyOf) Filter(filter types.IntervalsFilterVariant) *_intervalsAnyOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAnyOf) Intervals(intervals ...types.IntervalsVariant) *_intervalsAnyOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAnyOf) IntervalsValues(intervalsvalues []types.Intervals) *_intervalsAnyOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAnyOf) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervalsAnyOf) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAnyOf) IntervalsAnyOfCaster() *types.IntervalsAnyOf {
	_ = "STUB: not implemented"
	return nil
}
