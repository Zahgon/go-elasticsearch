package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _executionState struct {
	v *types.ExecutionState
}

func NewExecutionState(successful bool) *_executionState { _ = "STUB: not implemented"; return nil }

func (s *_executionState) Reason(reason string) *_executionState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_executionState) Successful(successful bool) *_executionState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_executionState) Timestamp(datetime types.DateTimeVariant) *_executionState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_executionState) ExecutionStateCaster() *types.ExecutionState {
	_ = "STUB: not implemented"
	return nil
}
