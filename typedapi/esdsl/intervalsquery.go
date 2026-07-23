package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsQuery struct {
	v *types.IntervalsQuery
}

func NewIntervalsQuery() *_intervalsQuery { _ = "STUB: not implemented"; return nil }

func (s *_intervalsQuery) AllOf(allof types.IntervalsAllOfVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) AnyOf(anyof types.IntervalsAnyOfVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Fuzzy(fuzzy types.IntervalsFuzzyVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Match(match types.IntervalsMatchVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Prefix(prefix types.IntervalsPrefixVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Range(range_ types.IntervalsRangeVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Regexp(regexp types.IntervalsRegexpVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Wildcard(wildcard types.IntervalsWildcardVariant) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) Boost(boost float32) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) QueryName_(queryname_ string) *_intervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsQuery) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}
