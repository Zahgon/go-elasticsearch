package types

type LinearInterpolationSmoothingModel struct {
	BigramLambda  Float64 `json:"bigram_lambda"`
	TrigramLambda Float64 `json:"trigram_lambda"`
	UnigramLambda Float64 `json:"unigram_lambda"`
}

func (s *LinearInterpolationSmoothingModel) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLinearInterpolationSmoothingModel() *LinearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

type LinearInterpolationSmoothingModelVariant interface {
	LinearInterpolationSmoothingModelCaster() *LinearInterpolationSmoothingModel
}

func (s *LinearInterpolationSmoothingModel) LinearInterpolationSmoothingModelCaster() *LinearInterpolationSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}
