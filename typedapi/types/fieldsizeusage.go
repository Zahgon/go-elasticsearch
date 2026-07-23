package types

type FieldSizeUsage struct {
	Size        ByteSize `json:"size,omitempty"`
	SizeInBytes int64    `json:"size_in_bytes"`
}

func (s *FieldSizeUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldSizeUsage() *FieldSizeUsage { _ = "STUB: not implemented"; return nil }
