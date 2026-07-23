package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _elserServiceSettings struct {
	v *types.ElserServiceSettings
}

func NewElserServiceSettings(numallocations int, numthreads int) *_elserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elserServiceSettings) AdaptiveAllocations(adaptiveallocations types.AdaptiveAllocationsVariant) *_elserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elserServiceSettings) NumAllocations(numallocations int) *_elserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elserServiceSettings) NumThreads(numthreads int) *_elserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_elserServiceSettings) ElserServiceSettingsCaster() *types.ElserServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
