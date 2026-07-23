package types

type SignificantStringTermsBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	BgCount      int64                `json:"bg_count"`
	DocCount     int64                `json:"doc_count"`
	Key          string               `json:"key"`
	Score        Float64              `json:"score"`
}

func (s *SignificantStringTermsBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SignificantStringTermsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSignificantStringTermsBucket() *SignificantStringTermsBucket {
	_ = "STUB: not implemented"
	return nil
}
