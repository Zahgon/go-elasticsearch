package types

type ExplanationDetail struct {
	Description string              `json:"description"`
	Details     []ExplanationDetail `json:"details,omitempty"`
	Value       float32             `json:"value"`
}

func (s *ExplanationDetail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExplanationDetail() *ExplanationDetail { _ = "STUB: not implemented"; return nil }
