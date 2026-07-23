package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _throttleState struct {
	v *types.ThrottleState
}

func NewThrottleState(reason string) *_throttleState { _ = "STUB: not implemented"; return nil }

func (s *_throttleState) Reason(reason string) *_throttleState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_throttleState) Timestamp(datetime types.DateTimeVariant) *_throttleState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_throttleState) ThrottleStateCaster() *types.ThrottleState {
	_ = "STUB: not implemented"
	return nil
}
