package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionop"
)

type WatcherCondition struct {
	Always       *AlwaysCondition                                  `json:"always,omitempty"`
	ArrayCompare map[string]ArrayCompareCondition                  `json:"array_compare,omitempty"`
	Compare      map[string]map[conditionop.ConditionOp]FieldValue `json:"compare,omitempty"`
	Never        *NeverCondition                                   `json:"never,omitempty"`
	Script       *ScriptCondition                                  `json:"script,omitempty"`
}

func NewWatcherCondition() *WatcherCondition { _ = "STUB: not implemented"; return nil }

type WatcherConditionVariant interface {
	WatcherConditionCaster() *WatcherCondition
}

func (s *WatcherCondition) WatcherConditionCaster() *WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}
