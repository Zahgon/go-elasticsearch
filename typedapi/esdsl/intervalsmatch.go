package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsMatch struct {
	v *types.IntervalsMatch
}

func NewIntervalsMatch(query string) *_intervalsMatch { _ = "STUB: not implemented"; return nil }

func (s *_intervalsMatch) Analyzer(analyzer string) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) Filter(filter types.IntervalsFilterVariant) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) MaxGaps(maxgaps int) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) Ordered(ordered bool) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) Query(query string) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) UseField(field string) *_intervalsMatch {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervalsMatch) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsMatch) IntervalsMatchCaster() *types.IntervalsMatch {
	_ = "STUB: not implemented"
	return nil
}
