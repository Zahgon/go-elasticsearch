package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsPrefix struct {
	v *types.IntervalsPrefix
}

func NewIntervalsPrefix(prefix string) *_intervalsPrefix { _ = "STUB: not implemented"; return nil }

func (s *_intervalsPrefix) Analyzer(analyzer string) *_intervalsPrefix {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsPrefix) Prefix(prefix string) *_intervalsPrefix {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsPrefix) UseField(field string) *_intervalsPrefix {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsPrefix) IntervalsCaster() *types.Intervals {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsPrefix) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsPrefix) IntervalsPrefixCaster() *types.IntervalsPrefix {
	_ = "STUB: not implemented"
	return nil
}
