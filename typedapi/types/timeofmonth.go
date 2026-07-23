package types

type TimeOfMonth struct {
	At []string `json:"at"`
	On []int    `json:"on"`
}

func NewTimeOfMonth() *TimeOfMonth { _ = "STUB: not implemented"; return nil }

type TimeOfMonthVariant interface {
	TimeOfMonthCaster() *TimeOfMonth
}

func (s *TimeOfMonth) TimeOfMonthCaster() *TimeOfMonth { _ = "STUB: not implemented"; return nil }
