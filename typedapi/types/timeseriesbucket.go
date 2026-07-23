package types

type TimeSeriesBucket struct {
	Aggregations map[string]Aggregate  `json:"-"`
	DocCount     int64                 `json:"doc_count"`
	Key          map[string]FieldValue `json:"key"`
}

func (s *TimeSeriesBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s TimeSeriesBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewTimeSeriesBucket() *TimeSeriesBucket { _ = "STUB: not implemented"; return nil }
