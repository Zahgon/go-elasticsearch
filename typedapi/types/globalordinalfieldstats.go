package types

type GlobalOrdinalFieldStats struct {
	BuildTime          *string `json:"build_time,omitempty"`
	BuildTimeInMillis  int64   `json:"build_time_in_millis"`
	ShardMaxValueCount int64   `json:"shard_max_value_count"`
}

func (s *GlobalOrdinalFieldStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewGlobalOrdinalFieldStats() *GlobalOrdinalFieldStats { _ = "STUB: not implemented"; return nil }
