package types

type IpRangeBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	From         *string              `json:"from,omitempty"`
	Key          *string              `json:"key,omitempty"`
	To           *string              `json:"to,omitempty"`
}

func (s *IpRangeBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IpRangeBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIpRangeBucket() *IpRangeBucket { _ = "STUB: not implemented"; return nil }
