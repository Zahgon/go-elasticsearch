package types

import (
	"encoding/json"
)

type AutoscalingPolicy struct {
	Deciders map[string]json.RawMessage `json:"deciders"`
	Roles    []string                   `json:"roles"`
}

func NewAutoscalingPolicy() *AutoscalingPolicy { _ = "STUB: not implemented"; return nil }

type AutoscalingPolicyVariant interface {
	AutoscalingPolicyCaster() *AutoscalingPolicy
}

func (s *AutoscalingPolicy) AutoscalingPolicyCaster() *AutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}
