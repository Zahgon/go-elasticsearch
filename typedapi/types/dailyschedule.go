package types

type DailySchedule struct {
	At []ScheduleTimeOfDay `json:"at"`
}

func (s *DailySchedule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDailySchedule() *DailySchedule { _ = "STUB: not implemented"; return nil }

type DailyScheduleVariant interface {
	DailyScheduleCaster() *DailySchedule
}

func (s *DailySchedule) DailyScheduleCaster() *DailySchedule { _ = "STUB: not implemented"; return nil }
