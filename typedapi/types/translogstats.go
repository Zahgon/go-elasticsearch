package types

type TranslogStats struct {
	EarliestLastModifiedAge int64   `json:"earliest_last_modified_age"`
	Operations              int64   `json:"operations"`
	Size                    *string `json:"size,omitempty"`
	SizeInBytes             int64   `json:"size_in_bytes"`
	UncommittedOperations   int     `json:"uncommitted_operations"`
	UncommittedSize         *string `json:"uncommitted_size,omitempty"`
	UncommittedSizeInBytes  int64   `json:"uncommitted_size_in_bytes"`
}

func (s *TranslogStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTranslogStats() *TranslogStats { _ = "STUB: not implemented"; return nil }
