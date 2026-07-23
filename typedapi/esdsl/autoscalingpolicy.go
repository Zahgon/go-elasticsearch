package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _autoscalingPolicy struct {
	v *types.AutoscalingPolicy
}

func NewAutoscalingPolicy() *_autoscalingPolicy { _ = "STUB: not implemented"; return nil }

func (s *_autoscalingPolicy) Deciders(deciders map[string]json.RawMessage) *_autoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoscalingPolicy) AddDecider(key string, value json.RawMessage) *_autoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoscalingPolicy) Roles(roles ...string) *_autoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoscalingPolicy) AutoscalingPolicyCaster() *types.AutoscalingPolicy {
	_ = "STUB: not implemented"
	return nil
}
