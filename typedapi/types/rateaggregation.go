package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/calendarinterval"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ratemode"
)

type RateAggregation struct {
	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Missing Missing `json:"missing,omitempty"`

	Mode   *ratemode.RateMode `json:"mode,omitempty"`
	Script *Script            `json:"script,omitempty"`

	Unit *calendarinterval.CalendarInterval `json:"unit,omitempty"`
}

func (s *RateAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRateAggregation() *RateAggregation { _ = "STUB: not implemented"; return nil }

type RateAggregationVariant interface {
	RateAggregationCaster() *RateAggregation
}

func (s *RateAggregation) RateAggregationCaster() *RateAggregation {
	_ = "STUB: not implemented"
	return nil
}
