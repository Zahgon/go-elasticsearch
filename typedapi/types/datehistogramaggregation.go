package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
)

type DateHistogramAggregation struct {
	CalendarInterval *calendarinterval.CalendarInterval `json:"calendar_interval,omitempty"`

	ExtendedBounds *ExtendedBoundsFieldDateMath `json:"extended_bounds,omitempty"`

	Field *string `json:"field,omitempty"`

	FixedInterval Duration `json:"fixed_interval,omitempty"`

	Format *string `json:"format,omitempty"`

	HardBounds *ExtendedBoundsFieldDateMath `json:"hard_bounds,omitempty"`
	Interval   Duration                     `json:"interval,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	MinDocCount *int `json:"min_doc_count,omitempty"`

	Missing DateTime `json:"missing,omitempty"`

	Offset Duration `json:"offset,omitempty"`

	Order  AggregateOrder             `json:"order,omitempty"`
	Params map[string]json.RawMessage `json:"params,omitempty"`
	Script *Script                    `json:"script,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *DateHistogramAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateHistogramAggregation() *DateHistogramAggregation { _ = "STUB: not implemented"; return nil }

type DateHistogramAggregationVariant interface {
	DateHistogramAggregationCaster() *DateHistogramAggregation
}

func (s *DateHistogramAggregation) DateHistogramAggregationCaster() *DateHistogramAggregation {
	_ = "STUB: not implemented"
	return nil
}
