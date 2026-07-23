package types

type StringRareTermsBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
}

func (s *StringRareTermsBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s StringRareTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewStringRareTermsBucket() *StringRareTermsBucket { _ = "STUB: not implemented"; return nil }
