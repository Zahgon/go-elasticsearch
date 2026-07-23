package types

type MappingStats struct {
	TotalCount                    int64    `json:"total_count"`
	TotalEstimatedOverhead        ByteSize `json:"total_estimated_overhead,omitempty"`
	TotalEstimatedOverheadInBytes int64    `json:"total_estimated_overhead_in_bytes"`
}

func (s *MappingStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMappingStats() *MappingStats { _ = "STUB: not implemented"; return nil }
