package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _adaptiveAllocationsSettings struct {
	v *types.AdaptiveAllocationsSettings
}

func NewAdaptiveAllocationsSettings(enabled bool) *_adaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocationsSettings) Enabled(enabled bool) *_adaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocationsSettings) MaxNumberOfAllocations(maxnumberofallocations int) *_adaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocationsSettings) MinNumberOfAllocations(minnumberofallocations int) *_adaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_adaptiveAllocationsSettings) AdaptiveAllocationsSettingsCaster() *types.AdaptiveAllocationsSettings {
	_ = "STUB: not implemented"
	return nil
}
