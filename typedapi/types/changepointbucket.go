package types

type ChangePointBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          FieldValue           `json:"key"`
}

func (s *ChangePointBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s ChangePointBucket) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewChangePointBucket() *ChangePointBucket { _ = "STUB: not implemented"; return nil }
