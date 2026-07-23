package types

type RangeBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	From         *Float64             `json:"from,omitempty"`
	FromAsString *string              `json:"from_as_string,omitempty"`

	Key        *string  `json:"key,omitempty"`
	To         *Float64 `json:"to,omitempty"`
	ToAsString *string  `json:"to_as_string,omitempty"`
}

func (s *RangeBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s RangeBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewRangeBucket() *RangeBucket { _ = "STUB: not implemented"; return nil }
