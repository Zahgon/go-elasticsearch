package types

type CCSUsageStats struct {
	Clients map[string]int `json:"clients"`

	Clusters map[string]CCSUsageClusterStats `json:"clusters"`

	FailureReasons map[string]int `json:"failure_reasons"`

	Features map[string]int `json:"features"`

	RemotesPerSearchAvg Float64 `json:"remotes_per_search_avg"`

	RemotesPerSearchMax int `json:"remotes_per_search_max"`

	Skipped int `json:"skipped"`

	Success int `json:"success"`

	Took CCSUsageTimeValue `json:"took"`

	TookMrtFalse *CCSUsageTimeValue `json:"took_mrt_false,omitempty"`

	TookMrtTrue *CCSUsageTimeValue `json:"took_mrt_true,omitempty"`

	Total int `json:"total"`
}

func (s *CCSUsageStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCCSUsageStats() *CCSUsageStats { _ = "STUB: not implemented"; return nil }
