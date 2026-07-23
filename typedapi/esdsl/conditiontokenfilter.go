package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _conditionTokenFilter struct {
	v *types.ConditionTokenFilter
}

func NewConditionTokenFilter(script types.ScriptVariant) *_conditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_conditionTokenFilter) Filter(filters ...string) *_conditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_conditionTokenFilter) Script(script types.ScriptVariant) *_conditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_conditionTokenFilter) Version(versionstring string) *_conditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_conditionTokenFilter) ConditionTokenFilterCaster() *types.ConditionTokenFilter {
	_ = "STUB: not implemented"
	return nil
}
