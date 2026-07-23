package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsWildcard struct {
	v *types.IntervalsWildcard
}

func NewIntervalsWildcard(pattern string) *_intervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) Analyzer(analyzer string) *_intervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) Pattern(pattern string) *_intervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) UseField(field string) *_intervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) IntervalsCaster() *types.Intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsWildcard) IntervalsWildcardCaster() *types.IntervalsWildcard {
	_ = "STUB: not implemented"
	return nil
}
