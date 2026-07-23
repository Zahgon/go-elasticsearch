package types

import (
	"encoding/json"
)

type AutoscalingDecider struct {
	ReasonDetails    json.RawMessage     `json:"reason_details,omitempty"`
	ReasonSummary    *string             `json:"reason_summary,omitempty"`
	RequiredCapacity AutoscalingCapacity `json:"required_capacity"`
}

func (s *AutoscalingDecider) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoscalingDecider() *AutoscalingDecider { _ = "STUB: not implemented"; return nil }
