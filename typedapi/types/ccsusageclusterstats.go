package types

type CCSUsageClusterStats struct {
	Skipped int `json:"skipped"`

	Took CCSUsageTimeValue `json:"took"`

	Total int `json:"total"`
}

func (s *CCSUsageClusterStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCCSUsageClusterStats() *CCSUsageClusterStats { _ = "STUB: not implemented"; return nil }
