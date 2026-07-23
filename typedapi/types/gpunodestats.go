package types

type GpuNodeStats struct {
	Enabled bool `json:"enabled"`

	IndexBuildCount int64 `json:"index_build_count"`

	MemoryInBytes int64 `json:"memory_in_bytes"`

	Type string `json:"type"`
}

func (s *GpuNodeStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGpuNodeStats() *GpuNodeStats { _ = "STUB: not implemented"; return nil }
