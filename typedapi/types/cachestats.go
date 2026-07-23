package types

type CacheStats struct {
	Count              int    `json:"count"`
	Evictions          int    `json:"evictions"`
	Hits               int    `json:"hits"`
	HitsTimeInMillis   int64  `json:"hits_time_in_millis"`
	Misses             int    `json:"misses"`
	MissesTimeInMillis int64  `json:"misses_time_in_millis"`
	NodeId             string `json:"node_id"`
	SizeInBytes        int64  `json:"size_in_bytes"`
}

func (s *CacheStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCacheStats() *CacheStats { _ = "STUB: not implemented"; return nil }
