package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _analysisLimits struct {
	v *types.AnalysisLimits
}

func NewAnalysisLimits() *_analysisLimits { _ = "STUB: not implemented"; return nil }

func (s *_analysisLimits) CategorizationExamplesLimit(categorizationexampleslimit int64) *_analysisLimits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisLimits) ModelMemoryLimit(bytesize types.ByteSizeVariant) *_analysisLimits {
	_ = "STUB: not implemented"
	return nil
}

func (s *_analysisLimits) AnalysisLimitsCaster() *types.AnalysisLimits {
	_ = "STUB: not implemented"
	return nil
}
