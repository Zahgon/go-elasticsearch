package types

type RollupJobSummaryField struct {
	Agg              string   `json:"agg"`
	CalendarInterval Duration `json:"calendar_interval,omitempty"`
	TimeZone         *string  `json:"time_zone,omitempty"`
}

func (s *RollupJobSummaryField) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRollupJobSummaryField() *RollupJobSummaryField { _ = "STUB: not implemented"; return nil }
