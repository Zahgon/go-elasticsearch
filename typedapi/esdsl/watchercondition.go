package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watcherCondition struct {
	v *types.WatcherCondition
}

func NewWatcherCondition() *_watcherCondition { _ = "STUB: not implemented"; return nil }

func (s *_watcherCondition) Always(always types.AlwaysConditionVariant) *_watcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherCondition) ArrayCompare(key string, value types.ArrayCompareConditionVariant) *_watcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherCondition) Never(never types.NeverConditionVariant) *_watcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherCondition) Script(script types.ScriptConditionVariant) *_watcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_watcherCondition) WatcherConditionCaster() *types.WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}
