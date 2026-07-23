package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

type _downsampleConfig struct {
	v *types.DownsampleConfig
}

func NewDownsampleConfig() *_downsampleConfig { _ = "STUB: not implemented"; return nil }

func (s *_downsampleConfig) FixedInterval(durationlarge string) *_downsampleConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsampleConfig) SamplingMethod(samplingmethod samplingmethod.SamplingMethod) *_downsampleConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsampleConfig) DownsampleConfigCaster() *types.DownsampleConfig {
	_ = "STUB: not implemented"
	return nil
}
