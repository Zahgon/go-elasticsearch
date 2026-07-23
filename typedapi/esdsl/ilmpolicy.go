package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ilmPolicy struct {
	v *types.IlmPolicy
}

func NewIlmPolicy(phases types.PhasesVariant) *_ilmPolicy { _ = "STUB: not implemented"; return nil }

func (s *_ilmPolicy) Meta_(metadata types.MetadataVariant) *_ilmPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmPolicy) Phases(phases types.PhasesVariant) *_ilmPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ilmPolicy) IlmPolicyCaster() *types.IlmPolicy { _ = "STUB: not implemented"; return nil }
