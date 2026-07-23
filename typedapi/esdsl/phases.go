package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _phases struct {
	v *types.Phases
}

func NewPhases() *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) Cold(cold types.PhaseVariant) *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) Delete(delete types.PhaseVariant) *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) Frozen(frozen types.PhaseVariant) *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) Hot(hot types.PhaseVariant) *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) Warm(warm types.PhaseVariant) *_phases { _ = "STUB: not implemented"; return nil }

func (s *_phases) PhasesCaster() *types.Phases { _ = "STUB: not implemented"; return nil }
