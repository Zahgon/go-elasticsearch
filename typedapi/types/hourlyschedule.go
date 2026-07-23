package types

type HourlySchedule struct {
	Minute []int `json:"minute"`
}

func NewHourlySchedule() *HourlySchedule { _ = "STUB: not implemented"; return nil }

type HourlyScheduleVariant interface {
	HourlyScheduleCaster() *HourlySchedule
}

func (s *HourlySchedule) HourlyScheduleCaster() *HourlySchedule {
	_ = "STUB: not implemented"
	return nil
}
