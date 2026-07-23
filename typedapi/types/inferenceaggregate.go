package types

import (
	"encoding/json"
)

type InferenceAggregate struct {
	Data              map[string]json.RawMessage   `json:"-"`
	FeatureImportance []InferenceFeatureImportance `json:"feature_importance,omitempty"`
	Meta              Metadata                     `json:"meta,omitempty"`
	TopClasses        []InferenceTopClassEntry     `json:"top_classes,omitempty"`
	Value             FieldValue                   `json:"value,omitempty"`
	Warning           *string                      `json:"warning,omitempty"`
}

func (s *InferenceAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s InferenceAggregate) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewInferenceAggregate() *InferenceAggregate { _ = "STUB: not implemented"; return nil }
