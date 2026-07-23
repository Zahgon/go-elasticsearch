package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _laplaceSmoothingModel struct {
	v *types.LaplaceSmoothingModel
}

func NewLaplaceSmoothingModel(alpha types.Float64) *_laplaceSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_laplaceSmoothingModel) Alpha(alpha types.Float64) *_laplaceSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_laplaceSmoothingModel) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_laplaceSmoothingModel) LaplaceSmoothingModelCaster() *types.LaplaceSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}
