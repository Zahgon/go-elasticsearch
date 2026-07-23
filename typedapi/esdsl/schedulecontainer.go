package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scheduleContainer struct {
	v *types.ScheduleContainer
}

func NewScheduleContainer() *_scheduleContainer { _ = "STUB: not implemented"; return nil }

func (s *_scheduleContainer) Cron(cronexpression string) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Daily(daily types.DailyScheduleVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Hourly(hourly types.HourlyScheduleVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Interval(duration types.DurationVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Monthly(monthlies ...types.TimeOfMonthVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Timezone(timezone string) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Weekly(weeklies ...types.TimeOfWeekVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) Yearly(yearlies ...types.TimeOfYearVariant) *_scheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleContainer) ScheduleContainerCaster() *types.ScheduleContainer {
	_ = "STUB: not implemented"
	return nil
}
