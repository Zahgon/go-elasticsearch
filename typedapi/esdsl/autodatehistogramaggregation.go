package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/minimuminterval"
)

type _autoDateHistogramAggregation struct {
	v *types.AutoDateHistogramAggregation
}

func NewAutoDateHistogramAggregation() *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Buckets(buckets int) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Field(field string) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Format(format string) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) MinimumInterval(minimuminterval minimuminterval.MinimumInterval) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Missing(datetime types.DateTimeVariant) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Offset(offset string) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Params(params map[string]json.RawMessage) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) AddParam(key string, value json.RawMessage) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) Script(script types.ScriptVariant) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) TimeZone(timezone string) *_autoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) AggregationsCaster() *types.Aggregations {
	_ = "STUB: not implemented"
	return nil
}

func (s *_autoDateHistogramAggregation) AutoDateHistogramAggregationCaster() *types.AutoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
