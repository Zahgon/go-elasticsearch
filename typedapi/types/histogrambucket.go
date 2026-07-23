package types

type HistogramBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	Key          Float64              `json:"key"`
	KeyAsString  *string              `json:"key_as_string,omitempty"`
}

func (s *HistogramBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s HistogramBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewHistogramBucket() *HistogramBucket { _ = "STUB: not implemented"; return nil }
