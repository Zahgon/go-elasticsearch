package types

type StupidBackoffSmoothingModel struct {
	Discount Float64 `json:"discount"`
}

func (s *StupidBackoffSmoothingModel) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewStupidBackoffSmoothingModel() *StupidBackoffSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}

type StupidBackoffSmoothingModelVariant interface {
	StupidBackoffSmoothingModelCaster() *StupidBackoffSmoothingModel
}

func (s *StupidBackoffSmoothingModel) StupidBackoffSmoothingModelCaster() *StupidBackoffSmoothingModel {
	_ = "STUB: not implemented"
	return nil
}
