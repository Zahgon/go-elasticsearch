package types

type ShardQueryCache struct {
	CacheCount        int64 `json:"cache_count"`
	CacheSize         int64 `json:"cache_size"`
	Evictions         int64 `json:"evictions"`
	HitCount          int64 `json:"hit_count"`
	MemorySizeInBytes int64 `json:"memory_size_in_bytes"`
	MissCount         int64 `json:"miss_count"`
	TotalCount        int64 `json:"total_count"`
}

func (s *ShardQueryCache) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardQueryCache() *ShardQueryCache { _ = "STUB: not implemented"; return nil }
