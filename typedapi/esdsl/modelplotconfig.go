package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _modelPlotConfig struct {
	v *types.ModelPlotConfig
}

func NewModelPlotConfig() *_modelPlotConfig { _ = "STUB: not implemented"; return nil }

func (s *_modelPlotConfig) AnnotationsEnabled(annotationsenabled bool) *_modelPlotConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_modelPlotConfig) Enabled(enabled bool) *_modelPlotConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_modelPlotConfig) Terms(field string) *_modelPlotConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_modelPlotConfig) ModelPlotConfigCaster() *types.ModelPlotConfig {
	_ = "STUB: not implemented"
	return nil
}
