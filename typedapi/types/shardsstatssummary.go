package types

type ShardsStatsSummary struct {
	Incremental       ShardsStatsSummaryItem `json:"incremental"`
	StartTimeInMillis int64                  `json:"start_time_in_millis"`
	Time              Duration               `json:"time,omitempty"`
	TimeInMillis      int64                  `json:"time_in_millis"`
	Total             ShardsStatsSummaryItem `json:"total"`
}

func (s *ShardsStatsSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewShardsStatsSummary() *ShardsStatsSummary { _ = "STUB: not implemented"; return nil }
