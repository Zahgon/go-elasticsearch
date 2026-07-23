package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _neverCondition struct {
	v *types.NeverCondition
}

func NewNeverCondition() *_neverCondition { _ = "STUB: not implemented"; return nil }

func (s *_neverCondition) WatcherConditionCaster() *types.WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_neverCondition) NeverConditionCaster() *types.NeverCondition {
	_ = "STUB: not implemented"
	return nil
}
