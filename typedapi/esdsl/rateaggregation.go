package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ratemode"
)

type _rateAggregation struct {
	v *types.RateAggregation
}

func NewRateAggregation() *_rateAggregation { _ = "STUB: not implemented"; return nil }

func (s *_rateAggregation) Mode(mode ratemode.RateMode) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) Unit(unit calendarinterval.CalendarInterval) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) Field(field string) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) Format(format string) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) Missing(missing types.MissingVariant) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) Script(script types.ScriptVariant) *_rateAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rateAggregation) RateAggregationCaster() *types.RateAggregation {
	_ = "STUB: not implemented"
	return nil
}
