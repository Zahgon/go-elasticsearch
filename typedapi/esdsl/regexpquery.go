package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _regexpQuery struct {
	k string
	v *types.RegexpQuery
}

func NewRegexpQuery(field string, value string) *_regexpQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexpQuery) CaseInsensitive(caseinsensitive bool) *_regexpQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexpQuery) Flags(flags string) *_regexpQuery { _ = "STUB: not implemented"; return nil }

func (s *_regexpQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_regexpQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexpQuery) Rewrite(multitermqueryrewrite string) *_regexpQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexpQuery) Value(value string) *_regexpQuery { _ = "STUB: not implemented"; return nil }

func (s *_regexpQuery) Boost(boost float32) *_regexpQuery { _ = "STUB: not implemented"; return nil }

func (s *_regexpQuery) QueryName_(queryname_ string) *_regexpQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexpQuery) QueryCaster() *types.Query { _ = "STUB: not implemented"; return nil }

func NewSingleRegexpQuery() *_regexpQuery { _ = "STUB: not implemented"; return nil }

func (s *_regexpQuery) RegexpQueryCaster() *types.RegexpQuery {
	_ = "STUB: not implemented"
	return nil
}
