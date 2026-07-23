package types

type LaplaceSmoothingModel struct {
	Alpha Float64 `json:"alpha"`
}

func (s *LaplaceSmoothingModel) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewLaplaceSmoothingModel() *LaplaceSmoothingModel { _ = "STUB: not implemented"; return nil }

type LaplaceSmoothingModelVariant interface {
	LaplaceSmoothingModelCaster() *LaplaceSmoothingModel
}

func (s *LaplaceSmoothingModel) LaplaceSmoothingModelCaster() *LaplaceSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}
