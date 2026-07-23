package types

type ClusterStatistics struct {
	Details    map[string]ClusterDetails `json:"details,omitempty"`
	Failed     int                       `json:"failed"`
	Partial    int                       `json:"partial"`
	Running    int                       `json:"running"`
	Skipped    int                       `json:"skipped"`
	Successful int                       `json:"successful"`
	Total      int                       `json:"total"`
}

func (s *ClusterStatistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterStatistics() *ClusterStatistics { _ = "STUB: not implemented"; return nil }
