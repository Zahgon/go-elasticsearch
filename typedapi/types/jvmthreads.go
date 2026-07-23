package types

type JvmThreads struct {
	Count *int64 `json:"count,omitempty"`

	PeakCount *int64 `json:"peak_count,omitempty"`
}

func (s *JvmThreads) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJvmThreads() *JvmThreads { _ = "STUB: not implemented"; return nil }
