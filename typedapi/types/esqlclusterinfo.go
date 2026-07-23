package types

type EsqlClusterInfo struct {
	Details    map[string]EsqlClusterDetails `json:"details"`
	Failed     int                           `json:"failed"`
	Partial    int                           `json:"partial"`
	Running    int                           `json:"running"`
	Skipped    int                           `json:"skipped"`
	Successful int                           `json:"successful"`
	Total      int                           `json:"total"`
}

func (s *EsqlClusterInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewEsqlClusterInfo() *EsqlClusterInfo { _ = "STUB: not implemented"; return nil }
