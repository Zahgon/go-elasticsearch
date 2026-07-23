package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsAllOf struct {
	v *types.IntervalsAllOf
}

func NewIntervalsAllOf() *_intervalsAllOf { _ = "STUB: not implemented"; return nil }

func (s *_intervalsAllOf) Filter(filter types.IntervalsFilterVariant) *_intervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) Intervals(intervals ...types.IntervalsVariant) *_intervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) IntervalsValues(intervalsvalues []types.Intervals) *_intervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) MaxGaps(maxgaps int) *_intervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) Ordered(ordered bool) *_intervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervalsAllOf) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsAllOf) IntervalsAllOfCaster() *types.IntervalsAllOf {
	_ = "STUB: not implemented"
	return nil
}
