package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsFuzzy struct {
	v *types.IntervalsFuzzy
}

func NewIntervalsFuzzy(term string) *_intervalsFuzzy { _ = "STUB: not implemented"; return nil }

func (s *_intervalsFuzzy) Analyzer(analyzer string) *_intervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) Fuzziness(fuzziness types.FuzzinessVariant) *_intervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) PrefixLength(prefixlength int) *_intervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) Term(term string) *_intervalsFuzzy { _ = "STUB: not implemented"; return nil }

func (s *_intervalsFuzzy) Transpositions(transpositions bool) *_intervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) UseField(field string) *_intervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervalsFuzzy) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFuzzy) IntervalsFuzzyCaster() *types.IntervalsFuzzy {
	_ = "STUB: not implemented"
	return nil
}
