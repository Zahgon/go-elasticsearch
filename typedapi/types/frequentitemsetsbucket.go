package types

type FrequentItemSetsBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          map[string][]string  `json:"key"`
	Support      Float64              `json:"support"`
}

func (s *FrequentItemSetsBucket) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s FrequentItemSetsBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewFrequentItemSetsBucket() *FrequentItemSetsBucket { _ = "STUB: not implemented"; return nil }
