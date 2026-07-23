package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _smoothingModelContainer struct {
	v *types.SmoothingModelContainer
}

func NewSmoothingModelContainer() *_smoothingModelContainer { _ = "STUB: not implemented"; return nil }

func (s *_smoothingModelContainer) Laplace(laplace types.LaplaceSmoothingModelVariant) *_smoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_smoothingModelContainer) LinearInterpolation(linearinterpolation types.LinearInterpolationSmoothingModelVariant) *_smoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_smoothingModelContainer) StupidBackoff(stupidbackoff types.StupidBackoffSmoothingModelVariant) *_smoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_smoothingModelContainer) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}
