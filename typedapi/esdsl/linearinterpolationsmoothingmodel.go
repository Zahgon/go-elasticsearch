package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _linearInterpolationSmoothingModel struct {
	v *types.LinearInterpolationSmoothingModel
}

func NewLinearInterpolationSmoothingModel(bigramlambda types.Float64, trigramlambda types.Float64, unigramlambda types.Float64) *_linearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearInterpolationSmoothingModel) BigramLambda(bigramlambda types.Float64) *_linearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearInterpolationSmoothingModel) TrigramLambda(trigramlambda types.Float64) *_linearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearInterpolationSmoothingModel) UnigramLambda(unigramlambda types.Float64) *_linearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearInterpolationSmoothingModel) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_linearInterpolationSmoothingModel) LinearInterpolationSmoothingModelCaster() *types.LinearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}
