package types

type InferenceClassImportance struct {
	ClassName  string  `json:"class_name"`
	Importance Float64 `json:"importance"`
}

func (s *InferenceClassImportance) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewInferenceClassImportance() *InferenceClassImportance { _ = "STUB: not implemented"; return nil }
