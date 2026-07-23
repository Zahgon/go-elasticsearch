package types

type DoubleTermsBucket struct {
	Aggregations            map[string]Aggregate `json:"-"`
	DocCount                int64                `json:"doc_count"`
	DocCountErrorUpperBound *int64               `json:"doc_count_error_upper_bound,omitempty"`
	Key                     Float64              `json:"key"`
	KeyAsString             *string              `json:"key_as_string,omitempty"`
}

func (s *DoubleTermsBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s DoubleTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewDoubleTermsBucket() *DoubleTermsBucket { _ = "STUB: not implemented"; return nil }
