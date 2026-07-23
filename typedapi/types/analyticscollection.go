package types

type AnalyticsCollection struct {
	EventDataStream EventDataStream `json:"event_data_stream"`
}

func NewAnalyticsCollection() *AnalyticsCollection { _ = "STUB: not implemented"; return nil }
