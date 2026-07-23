package types

type LongTermsBucket struct {
	Aggregations            map[string]Aggregate `json:"-"`
	DocCount                int64                `json:"doc_count"`
	DocCountErrorUpperBound *int64               `json:"doc_count_error_upper_bound,omitempty"`
	Key                     int64                `json:"key"`
	KeyAsString             *string              `json:"key_as_string,omitempty"`
}

func (s *LongTermsBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s LongTermsBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewLongTermsBucket() *LongTermsBucket { _ = "STUB: not implemented"; return nil }
