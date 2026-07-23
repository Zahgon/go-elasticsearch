package types

type Explanation struct {
	Description string              `json:"description"`
	Details     []ExplanationDetail `json:"details"`
	Value       float32             `json:"value"`
}

func (s *Explanation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExplanation() *Explanation { _ = "STUB: not implemented"; return nil }
