package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _analysisConfig struct {
	v *types.AnalysisConfig
}

func NewAnalysisConfig() *_analysisConfig { _ = "STUB: not implemented"; return nil }

func (s *_analysisConfig) BucketSpan(duration types.DurationVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) CategorizationAnalyzer(categorizationanalyzer types.CategorizationAnalyzerVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) CategorizationFieldName(field string) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) CategorizationFilters(categorizationfilters ...string) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) Detectors(detectors ...types.DetectorVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) DetectorsValues(detectorsvalues []types.Detector) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) Influencers(influencers ...string) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) Latency(duration types.DurationVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) ModelPruneWindow(duration types.DurationVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) MultivariateByFields(multivariatebyfields bool) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) PerPartitionCategorization(perpartitioncategorization types.PerPartitionCategorizationVariant) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) SummaryCountFieldName(field string) *_analysisConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisConfig) AnalysisConfigCaster() *types.AnalysisConfig {
	_ = "STUB: not implemented"
	return nil
}
