package types

type Analytics struct {
	Available bool                `json:"available"`
	Enabled   bool                `json:"enabled"`
	Stats     AnalyticsStatistics `json:"stats"`
}

func (s *Analytics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalytics() *Analytics { _ = "STUB: not implemented"; return nil }
