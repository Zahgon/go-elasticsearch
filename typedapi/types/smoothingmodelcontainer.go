package types

type SmoothingModelContainer struct {
	Laplace *LaplaceSmoothingModel `json:"laplace,omitempty"`

	LinearInterpolation *LinearInterpolationSmoothingModel `json:"linear_interpolation,omitempty"`

	StupidBackoff *StupidBackoffSmoothingModel `json:"stupid_backoff,omitempty"`
}

func NewSmoothingModelContainer() *SmoothingModelContainer { _ = "STUB: not implemented"; return nil }

type SmoothingModelContainerVariant interface {
	SmoothingModelContainerCaster() *SmoothingModelContainer
}

func (s *SmoothingModelContainer) SmoothingModelContainerCaster() *SmoothingModelContainer {
	_ = "STUB: not implemented"
	return nil
}
