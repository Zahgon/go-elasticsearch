package types

type IndexSettingsTimeSeries struct {
	EndTime   DateTime `json:"end_time,omitempty"`
	StartTime DateTime `json:"start_time,omitempty"`
}

func (s *IndexSettingsTimeSeries) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingsTimeSeries() *IndexSettingsTimeSeries { _ = "STUB: not implemented"; return nil }

type IndexSettingsTimeSeriesVariant interface {
	IndexSettingsTimeSeriesCaster() *IndexSettingsTimeSeries
}

func (s *IndexSettingsTimeSeries) IndexSettingsTimeSeriesCaster() *IndexSettingsTimeSeries {
	_ = "STUB: not implemented"
	return nil
}
