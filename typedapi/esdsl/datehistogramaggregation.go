package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
)

type _dateHistogramAggregation struct {
	v *types.DateHistogramAggregation
}

func NewDateHistogramAggregation() *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) CalendarInterval(calendarinterval calendarinterval.CalendarInterval) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) ExtendedBounds(extendedbounds types.ExtendedBoundsFieldDateMathVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Field(field string) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) FixedInterval(duration types.DurationVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Format(format string) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) HardBounds(hardbounds types.ExtendedBoundsFieldDateMathVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Interval(duration types.DurationVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Keyed(keyed bool) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) MinDocCount(mindoccount int) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Missing(datetime types.DateTimeVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Offset(duration types.DurationVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Order(aggregateorder types.AggregateOrderVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Params(params map[string]json.RawMessage) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) AddParam(key string, value json.RawMessage) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) Script(script types.ScriptVariant) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) TimeZone(timezone string) *_dateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateHistogramAggregation) DateHistogramAggregationCaster() *types.DateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
