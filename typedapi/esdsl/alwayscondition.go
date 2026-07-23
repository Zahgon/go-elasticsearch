package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _alwaysCondition struct {
	v *types.AlwaysCondition
}

func NewAlwaysCondition() *_alwaysCondition { _ = "STUB: not implemented"; return nil }

func (s *_alwaysCondition) WatcherConditionCaster() *types.WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_alwaysCondition) AlwaysConditionCaster() *types.AlwaysCondition {
	_ = "STUB: not implemented"
	return nil
}
