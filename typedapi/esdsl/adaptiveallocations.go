package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _adaptiveAllocations struct {
	v *types.AdaptiveAllocations
}

func NewAdaptiveAllocations() *_adaptiveAllocations { _ = "STUB: not implemented"; return nil }

func (s *_adaptiveAllocations) Enabled(enabled bool) *_adaptiveAllocations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocations) MaxNumberOfAllocations(maxnumberofallocations int) *_adaptiveAllocations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocations) MinNumberOfAllocations(minnumberofallocations int) *_adaptiveAllocations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocations) AdaptiveAllocationsCaster() *types.AdaptiveAllocations {
	_ = "STUB: not implemented"
	return nil
}
