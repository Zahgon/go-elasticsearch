package types

type MultiTermsBucket struct {
	Aggregations            map[string]Aggregate `json:"-"`
	DocCount                int64                `json:"doc_count"`
	DocCountErrorUpperBound *int64               `json:"doc_count_error_upper_bound,omitempty"`
	Key                     []FieldValue         `json:"key"`
	KeyAsString             *string              `json:"key_as_string,omitempty"`
}

func (s *MultiTermsBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s MultiTermsBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewMultiTermsBucket() *MultiTermsBucket { _ = "STUB: not implemented"; return nil }
