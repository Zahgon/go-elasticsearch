package types

type LongRareTermsBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          int64                `json:"key"`
	KeyAsString  *string              `json:"key_as_string,omitempty"`
}

func (s *LongRareTermsBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s LongRareTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewLongRareTermsBucket() *LongRareTermsBucket { _ = "STUB: not implemented"; return nil }
