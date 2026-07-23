package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/minimuminterval"
)

type AutoDateHistogramAggregation struct {
	Buckets *int `json:"buckets,omitempty"`

	Field *string `json:"field,omitempty"`

	Format *string `json:"format,omitempty"`

	MinimumInterval *minimuminterval.MinimumInterval `json:"minimum_interval,omitempty"`

	Missing DateTime `json:"missing,omitempty"`

	Offset *string                    `json:"offset,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Script *Script                    `json:"script,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *AutoDateHistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAutoDateHistogramAggregation() *AutoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}

type AutoDateHistogramAggregationVariant interface {
	AutoDateHistogramAggregationCaster() *AutoDateHistogramAggregation
}

func (s *AutoDateHistogramAggregation) AutoDateHistogramAggregationCaster() *AutoDateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
