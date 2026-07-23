package types

type ValidationLoss struct {
	FoldValues []string `json:"fold_values"`

	LossType string `json:"loss_type"`
}

func (s *ValidationLoss) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewValidationLoss() *ValidationLoss { _ = "STUB: not implemented"; return nil }
