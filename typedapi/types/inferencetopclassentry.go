package types

type InferenceTopClassEntry struct {
	ClassName        FieldValue `json:"class_name"`
	ClassProbability Float64    `json:"class_probability"`
	ClassScore       Float64    `json:"class_score"`
}

func (s *InferenceTopClassEntry) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceTopClassEntry() *InferenceTopClassEntry { _ = "STUB: not implemented"; return nil }
