package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stepKey struct {
	v *types.StepKey
}

func NewStepKey(phase string) *_stepKey { _ = "STUB: not implemented"; return nil }

func (s *_stepKey) Action(action string) *_stepKey { _ = "STUB: not implemented"; return nil }

func (s *_stepKey) Name(name string) *_stepKey { _ = "STUB: not implemented"; return nil }

func (s *_stepKey) Phase(phase string) *_stepKey { _ = "STUB: not implemented"; return nil }

func (s *_stepKey) StepKeyCaster() *types.StepKey { _ = "STUB: not implemented"; return nil }
