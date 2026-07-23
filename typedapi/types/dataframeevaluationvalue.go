package types

type DataframeEvaluationValue struct {
	Value Float64 `json:"value"`
}

func (s *DataframeEvaluationValue) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeEvaluationValue() *DataframeEvaluationValue { _ = "STUB: not implemented"; return nil }
