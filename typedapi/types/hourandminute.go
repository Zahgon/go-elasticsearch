package types

type HourAndMinute struct {
	Hour   []int `json:"hour"`
	Minute []int `json:"minute"`
}

func NewHourAndMinute() *HourAndMinute { _ = "STUB: not implemented"; return nil }

type HourAndMinuteVariant interface {
	HourAndMinuteCaster() *HourAndMinute
}

func (s *HourAndMinute) HourAndMinuteCaster() *HourAndMinute { _ = "STUB: not implemented"; return nil }

func (s *HourAndMinute) ScheduleTimeOfDayCaster() *ScheduleTimeOfDay {
	_ = "STUB: not implemented"
	return nil
}
