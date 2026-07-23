package types

type ConfusionMatrixThreshold struct {
	FalseNegative int `json:"fn"`

	FalsePositive int `json:"fp"`

	TrueNegative int `json:"tn"`

	TruePositive int `json:"tp"`
}

func (s *ConfusionMatrixThreshold) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewConfusionMatrixThreshold() *ConfusionMatrixThreshold { _ = "STUB: not implemented"; return nil }
