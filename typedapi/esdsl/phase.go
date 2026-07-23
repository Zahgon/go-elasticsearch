package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _phase struct {
	v *types.Phase
}

func NewPhase() *_phase { _ = "STUB: not implemented"; return nil }

func (s *_phase) Actions(actions types.IlmActionsVariant) *_phase {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phase) MinAge(duration types.DurationVariant) *_phase {
	_ = "STUB: not implemented"
	return nil
}

func (s *_phase) PhaseCaster() *types.Phase { _ = "STUB: not implemented"; return nil }
