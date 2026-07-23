package types

type ScheduleContainer struct {
	Cron     *string         `json:"cron,omitempty"`
	Daily    *DailySchedule  `json:"daily,omitempty"`
	Hourly   *HourlySchedule `json:"hourly,omitempty"`
	Interval Duration        `json:"interval,omitempty"`
	Monthly  []TimeOfMonth   `json:"monthly,omitempty"`
	Timezone *string         `json:"timezone,omitempty"`
	Weekly   []TimeOfWeek    `json:"weekly,omitempty"`
	Yearly   []TimeOfYear    `json:"yearly,omitempty"`
}

func (s *ScheduleContainer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScheduleContainer() *ScheduleContainer { _ = "STUB: not implemented"; return nil }

type ScheduleContainerVariant interface {
	ScheduleContainerCaster() *ScheduleContainer
}

func (s *ScheduleContainer) ScheduleContainerCaster() *ScheduleContainer {
	_ = "STUB: not implemented"
	return nil
}
