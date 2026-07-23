package types

type WeightedAverageValue struct {
	Field *string `json:"field,omitempty"`

	Missing *Float64 `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

func (s *WeightedAverageValue) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewWeightedAverageValue() *WeightedAverageValue { _ = "STUB: not implemented"; return nil }

type WeightedAverageValueVariant interface {
	WeightedAverageValueCaster() *WeightedAverageValue
}

func (s *WeightedAverageValue) WeightedAverageValueCaster() *WeightedAverageValue {
	_ = "STUB: not implemented"
	return nil
}
