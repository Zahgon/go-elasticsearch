package types

type RollupFieldSummary struct {
	Agg              string   `json:"agg"`
	CalendarInterval Duration `json:"calendar_interval,omitempty"`
	TimeZone         *string  `json:"time_zone,omitempty"`
}

func (s *RollupFieldSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRollupFieldSummary() *RollupFieldSummary { _ = "STUB: not implemented"; return nil }
