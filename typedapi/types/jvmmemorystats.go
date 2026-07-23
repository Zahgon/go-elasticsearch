package types

type JvmMemoryStats struct {
	HeapCommittedInBytes *int64 `json:"heap_committed_in_bytes,omitempty"`

	HeapMax ByteSize `json:"heap_max,omitempty"`

	HeapMaxInBytes *int64 `json:"heap_max_in_bytes,omitempty"`

	HeapUsedInBytes *int64 `json:"heap_used_in_bytes,omitempty"`

	HeapUsedPercent *int64 `json:"heap_used_percent,omitempty"`

	NonHeapCommittedInBytes *int64 `json:"non_heap_committed_in_bytes,omitempty"`

	NonHeapUsedInBytes *int64 `json:"non_heap_used_in_bytes,omitempty"`

	Pools map[string]Pool `json:"pools,omitempty"`
}

func (s *JvmMemoryStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJvmMemoryStats() *JvmMemoryStats { _ = "STUB: not implemented"; return nil }
