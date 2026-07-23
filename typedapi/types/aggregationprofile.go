package types

type AggregationProfile struct {
	Breakdown   AggregationBreakdown     `json:"breakdown"`
	Children    []AggregationProfile     `json:"children,omitempty"`
	Debug       *AggregationProfileDebug `json:"debug,omitempty"`
	Description string                   `json:"description"`
	TimeInNanos int64                    `json:"time_in_nanos"`
	Type        string                   `json:"type"`
}

func (s *AggregationProfile) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAggregationProfile() *AggregationProfile { _ = "STUB: not implemented"; return nil }
