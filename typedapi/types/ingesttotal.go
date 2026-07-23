package types

type IngestTotal struct {
	Count int64 `json:"count"`

	Current int64 `json:"current"`

	Failed int64 `json:"failed"`

	TimeInMillis int64 `json:"time_in_millis"`
}

func (s *IngestTotal) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIngestTotal() *IngestTotal { _ = "STUB: not implemented"; return nil }
