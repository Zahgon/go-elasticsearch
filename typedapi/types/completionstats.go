package types

type CompletionStats struct {
	Fields map[string]FieldSizeUsage `json:"fields,omitempty"`

	Size ByteSize `json:"size,omitempty"`

	SizeInBytes int64 `json:"size_in_bytes"`
}

func (s *CompletionStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompletionStats() *CompletionStats { _ = "STUB: not implemented"; return nil }
