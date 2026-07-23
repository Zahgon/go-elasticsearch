package types

type DataStreamsStatsItem struct {
	BackingIndices int `json:"backing_indices"`

	DataStream string `json:"data_stream"`

	MaximumTimestamp int64 `json:"maximum_timestamp"`

	StoreSize ByteSize `json:"store_size,omitempty"`

	StoreSizeBytes int64 `json:"store_size_bytes"`
}

func (s *DataStreamsStatsItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamsStatsItem() *DataStreamsStatsItem { _ = "STUB: not implemented"; return nil }
