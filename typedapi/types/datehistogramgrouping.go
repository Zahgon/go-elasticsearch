package types

type DateHistogramGrouping struct {
	CalendarInterval Duration `json:"calendar_interval,omitempty"`

	Delay Duration `json:"delay,omitempty"`

	Field string `json:"field"`

	FixedInterval Duration `json:"fixed_interval,omitempty"`
	Format        *string  `json:"format,omitempty"`
	Interval      Duration `json:"interval,omitempty"`

	TimeZone *string `json:"time_zone,omitempty"`
}

func (s *DateHistogramGrouping) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDateHistogramGrouping() *DateHistogramGrouping { _ = "STUB: not implemented"; return nil }

type DateHistogramGroupingVariant interface {
	DateHistogramGroupingCaster() *DateHistogramGrouping
}

func (s *DateHistogramGrouping) DateHistogramGroupingCaster() *DateHistogramGrouping {
	_ = "STUB: not implemented"
	return nil
}
