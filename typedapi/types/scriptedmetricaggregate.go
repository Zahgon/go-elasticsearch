package types

import (
	"encoding/json"
)

type ScriptedMetricAggregate struct {
	Meta  Metadata        `json:"meta,omitempty"`
	Value json.RawMessage `json:"value,omitempty"`
}

func (s *ScriptedMetricAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewScriptedMetricAggregate() *ScriptedMetricAggregate { _ = "STUB: not implemented"; return nil }
