package types

type RequestCacheStats struct {
	Evictions         int64   `json:"evictions"`
	HitCount          int64   `json:"hit_count"`
	MemorySize        *string `json:"memory_size,omitempty"`
	MemorySizeInBytes int64   `json:"memory_size_in_bytes"`
	MissCount         int64   `json:"miss_count"`
}

func (s *RequestCacheStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRequestCacheStats() *RequestCacheStats { _ = "STUB: not implemented"; return nil }
