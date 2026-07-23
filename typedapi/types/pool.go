package types

type Pool struct {
	MaxInBytes *int64 `json:"max_in_bytes,omitempty"`

	PeakMaxInBytes *int64 `json:"peak_max_in_bytes,omitempty"`

	PeakUsedInBytes *int64 `json:"peak_used_in_bytes,omitempty"`

	UsedInBytes *int64 `json:"used_in_bytes,omitempty"`
}

func (s *Pool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPool() *Pool { _ = "STUB: not implemented"; return nil }
