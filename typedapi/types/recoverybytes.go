package types

type RecoveryBytes struct {
	Percent                      Percentage `json:"percent"`
	Recovered                    ByteSize   `json:"recovered,omitempty"`
	RecoveredFromSnapshot        ByteSize   `json:"recovered_from_snapshot,omitempty"`
	RecoveredFromSnapshotInBytes ByteSize   `json:"recovered_from_snapshot_in_bytes,omitempty"`
	RecoveredInBytes             ByteSize   `json:"recovered_in_bytes"`
	Reused                       ByteSize   `json:"reused,omitempty"`
	ReusedInBytes                ByteSize   `json:"reused_in_bytes"`
	Total                        ByteSize   `json:"total,omitempty"`
	TotalInBytes                 ByteSize   `json:"total_in_bytes"`
}

func (s *RecoveryBytes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRecoveryBytes() *RecoveryBytes { _ = "STUB: not implemented"; return nil }
