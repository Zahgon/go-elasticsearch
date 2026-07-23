package types

type DataframeEvaluationClass struct {
	ClassName string  `json:"class_name"`
	Value     Float64 `json:"value"`
}

func (s *DataframeEvaluationClass) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationClass() *DataframeEvaluationClass { _ = "STUB: not implemented"; return nil }
