package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _activationState struct {
	v *types.ActivationState
}

func NewActivationState(active bool) *_activationState { _ = "STUB: not implemented"; return nil }

func (s *_activationState) Active(active bool) *_activationState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_activationState) Timestamp(datetime types.DateTimeVariant) *_activationState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_activationState) ActivationStateCaster() *types.ActivationState {
	_ = "STUB: not implemented"
	return nil
}
