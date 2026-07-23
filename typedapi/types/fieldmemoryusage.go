package types

type FieldMemoryUsage struct {
	MemorySize        ByteSize `json:"memory_size,omitempty"`
	MemorySizeInBytes int64    `json:"memory_size_in_bytes"`
}

func (s *FieldMemoryUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldMemoryUsage() *FieldMemoryUsage { _ = "STUB: not implemented"; return nil }
