package types

type GarbageCollectorTotal struct {
	CollectionCount *int64 `json:"collection_count,omitempty"`

	CollectionTime *string `json:"collection_time,omitempty"`

	CollectionTimeInMillis *int64 `json:"collection_time_in_millis,omitempty"`
}

func (s *GarbageCollectorTotal) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGarbageCollectorTotal() *GarbageCollectorTotal { _ = "STUB: not implemented"; return nil }
