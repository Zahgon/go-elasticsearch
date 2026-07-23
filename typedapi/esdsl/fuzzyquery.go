package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _fuzzyQuery struct {
	k string
	v *types.FuzzyQuery
}

func NewFuzzyQuery(field string, value string) *_fuzzyQuery { _ = "STUB: not implemented"; return nil }

func (s *_fuzzyQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) MaxExpansions(maxexpansions int) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) PrefixLength(prefixlength int) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) Rewrite(multitermqueryrewrite string) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) Transpositions(transpositions bool) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) Value(value string) *_fuzzyQuery { _ = "STUB: not implemented"; return nil }

func (s *_fuzzyQuery) Boost(boost float32) *_fuzzyQuery { _ = "STUB: not implemented"; return nil }

func (s *_fuzzyQuery) QueryName_(queryname_ string) *_fuzzyQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fuzzyQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleFuzzyQuery() *_fuzzyQuery { _ = "STUB: not implemented"; return nil }

func (s *_fuzzyQuery) FuzzyQueryCaster() *types.FuzzyQuery { _ = "STUB: not implemented"; return nil }
