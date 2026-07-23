package types

type VariableWidthHistogramBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          Float64              `json:"key"`
	KeyAsString  *string              `json:"key_as_string,omitempty"`
	Max          Float64              `json:"max"`
	MaxAsString  *string              `json:"max_as_string,omitempty"`
	Min          Float64              `json:"min"`
	MinAsString  *string              `json:"min_as_string,omitempty"`
}

func (s *VariableWidthHistogramBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s VariableWidthHistogramBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewVariableWidthHistogramBucket() *VariableWidthHistogramBucket {
	_ = "STUB: not implemented"
	return nil
}
