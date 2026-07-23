package types

import (
	"encoding/json"
)

type ScriptedMetricAggregation struct {
	CombineScript *Script `json:"combine_script,omitempty"`

	Field *string `json:"field,omitempty"`

	InitScript *Script `json:"init_script,omitempty"`

	MapScript *Script `json:"map_script,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Params map[string]json.RawMessage `json:"params,omitempty"`

	ReduceScript *Script `json:"reduce_script,omitempty"`
	Script       *Script `json:"script,omitempty"`
}

func (s *ScriptedMetricAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewScriptedMetricAggregation() *ScriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}

type ScriptedMetricAggregationVariant interface {
	ScriptedMetricAggregationCaster() *ScriptedMetricAggregation
}

func (s *ScriptedMetricAggregation) ScriptedMetricAggregationCaster() *ScriptedMetricAggregation {
	_ = "STUB: not implemented"
	return nil
}
