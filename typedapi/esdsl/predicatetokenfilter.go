package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _predicateTokenFilter struct {
	v *types.PredicateTokenFilter
}

func NewPredicateTokenFilter(script types.ScriptVariant) *_predicateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_predicateTokenFilter) Script(script types.ScriptVariant) *_predicateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_predicateTokenFilter) Version(versionstring string) *_predicateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_predicateTokenFilter) PredicateTokenFilterCaster() *types.PredicateTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
