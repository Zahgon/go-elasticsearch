package types

type StringTermsBucket struct {
	Aggregations            map[string]Aggregate `json:"-"`
	DocCount                int64                `json:"doc_count"`
	DocCountErrorUpperBound *int64               `json:"doc_count_error_upper_bound,omitempty"`
	Key                     FieldValue           `json:"key"`
}

func (s *StringTermsBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s StringTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStringTermsBucket() *StringTermsBucket { _ = "STUB: not implemented"; return nil }
