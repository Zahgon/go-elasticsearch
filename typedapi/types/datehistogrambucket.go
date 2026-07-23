package types

type DateHistogramBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          int64                `json:"key"`
	KeyAsString  *string              `json:"key_as_string,omitempty"`
}

func (s *DateHistogramBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s DateHistogramBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDateHistogramBucket() *DateHistogramBucket { _ = "STUB: not implemented"; return nil }
