package types

type FiltersBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          *string              `json:"key,omitempty"`
}

func (s *FiltersBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s FiltersBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewFiltersBucket() *FiltersBucket { _ = "STUB: not implemented"; return nil }
