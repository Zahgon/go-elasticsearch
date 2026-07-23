package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _scriptedMetricAggregation struct {
	v *types.ScriptedMetricAggregation
}

func NewScriptedMetricAggregation() *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) CombineScript(combinescript types.ScriptVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) InitScript(initscript types.ScriptVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) MapScript(mapscript types.ScriptVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) Params(params map[string]json.RawMessage) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) AddParam(key string, value json.RawMessage) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) ReduceScript(reducescript types.ScriptVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) Field(field string) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) Missing(missing types.MissingVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) Script(script types.ScriptVariant) *_scriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptedMetricAggregation) ScriptedMetricAggregationCaster() *types.ScriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}
