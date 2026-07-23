package types

type ClusterJvmMemory struct {
	HeapMax ByteSize `json:"heap_max,omitempty"`

	HeapMaxInBytes int64 `json:"heap_max_in_bytes"`

	HeapUsed ByteSize `json:"heap_used,omitempty"`

	HeapUsedInBytes int64 `json:"heap_used_in_bytes"`
}

func (s *ClusterJvmMemory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterJvmMemory() *ClusterJvmMemory { _ = "STUB: not implemented"; return nil }
