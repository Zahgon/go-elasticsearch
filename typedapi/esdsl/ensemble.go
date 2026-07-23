package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ensemble struct {
	v *types.Ensemble
}

func NewEnsemble() *_ensemble { _ = "STUB: not implemented"; return nil }

func (s *_ensemble) AggregateOutput(aggregateoutput types.AggregateOutputVariant) *_ensemble {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ensemble) ClassificationLabels(classificationlabels ...string) *_ensemble {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ensemble) FeatureNames(featurenames ...string) *_ensemble {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ensemble) TargetType(targettype string) *_ensemble { _ = "STUB: not implemented"; return nil }

func (s *_ensemble) TrainedModels(trainedmodels ...types.TrainedModelVariant) *_ensemble {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ensemble) TrainedModelsValues(trainedmodelsvalues []types.TrainedModel) *_ensemble {
	_ = "STUB: not implemented"
	return nil
}

func (s *_ensemble) EnsembleCaster() *types.Ensemble { _ = "STUB: not implemented"; return nil }
