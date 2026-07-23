package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _delayedDataCheckConfig struct {
	v *types.DelayedDataCheckConfig
}

func NewDelayedDataCheckConfig(enabled bool) *_delayedDataCheckConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delayedDataCheckConfig) CheckWindow(duration types.DurationVariant) *_delayedDataCheckConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delayedDataCheckConfig) Enabled(enabled bool) *_delayedDataCheckConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_delayedDataCheckConfig) DelayedDataCheckConfigCaster() *types.DelayedDataCheckConfig {
	_ = "STUB: not implemented"
	return nil
}
