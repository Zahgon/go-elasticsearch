package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataframeAnalysisRegression struct {
	v *types.DataframeAnalysisRegression
}

func NewDataframeAnalysisRegression(dependentvariable string) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) LossFunction(lossfunction string) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) LossFunctionParameter(lossfunctionparameter types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) Alpha(alpha types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) DependentVariable(dependentvariable string) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) DownsampleFactor(downsamplefactor types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) EarlyStoppingEnabled(earlystoppingenabled bool) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) Eta(eta types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) EtaGrowthRatePerTree(etagrowthratepertree types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) FeatureBagFraction(featurebagfraction types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) FeatureProcessors(featureprocessors ...types.DataframeAnalysisFeatureProcessorVariant) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) FeatureProcessorsValues(featureprocessorsvalues []types.DataframeAnalysisFeatureProcessor) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) Gamma(gamma types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) Lambda(lambda types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) MaxOptimizationRoundsPerHyperparameter(maxoptimizationroundsperhyperparameter int) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) MaxTrees(maxtrees int) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) PredictionFieldName(field string) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) RandomizeSeed(randomizeseed types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) SoftTreeDepthLimit(softtreedepthlimit int) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) SoftTreeDepthTolerance(softtreedepthtolerance types.Float64) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) TrainingPercent(percentage types.PercentageVariant) *_dataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) DataframeAnalysisContainerCaster() *types.DataframeAnalysisContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataframeAnalysisRegression) DataframeAnalysisRegressionCaster() *types.DataframeAnalysisRegression {
	_ = "STUB: not implemented"
	return nil
}
