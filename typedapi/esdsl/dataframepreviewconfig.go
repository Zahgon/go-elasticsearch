package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframePreviewConfig struct {
	v *types.DataframePreviewConfig
}

func NewDataframePreviewConfig(analysis types.DataframeAnalysisContainerVariant, source types.DataframeAnalyticsSourceVariant) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) Analysis(analysis types.DataframeAnalysisContainerVariant) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) AnalyzedFields(analyzedfields types.DataframeAnalysisAnalyzedFieldsVariant) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) MaxNumThreads(maxnumthreads int) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) ModelMemoryLimit(modelmemorylimit string) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) Source(source types.DataframeAnalyticsSourceVariant) *_dataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframePreviewConfig) DataframePreviewConfigCaster() *types.DataframePreviewConfig {
	_ = "STUB: not implemented"
	return nil
}
