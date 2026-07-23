package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexSettingsTimeSeries struct {
	v *types.IndexSettingsTimeSeries
}

func NewIndexSettingsTimeSeries() *_indexSettingsTimeSeries { _ = "STUB: not implemented"; return nil }

func (s *_indexSettingsTimeSeries) EndTime(datetime types.DateTimeVariant) *_indexSettingsTimeSeries {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsTimeSeries) StartTime(datetime types.DateTimeVariant) *_indexSettingsTimeSeries {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexSettingsTimeSeries) IndexSettingsTimeSeriesCaster() *types.IndexSettingsTimeSeries {
	_ = "STUB: not implemented"
	return nil
}
