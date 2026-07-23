package types

type SignificantLongTermsBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	BgCount      int64                `json:"bg_count"`
	DocCount     int64                `json:"doc_count"`
	Key          int64                `json:"key"`
	KeyAsString  *string              `json:"key_as_string,omitempty"`
	Score        Float64              `json:"score"`
}

func (s *SignificantLongTermsBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SignificantLongTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSignificantLongTermsBucket() *SignificantLongTermsBucket {
	_ = "STUB: not implemented"
	return nil
}
