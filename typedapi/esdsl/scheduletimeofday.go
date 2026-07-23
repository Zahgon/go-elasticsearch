package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scheduleTimeOfDay struct {
	v types.ScheduleTimeOfDay
}

func NewScheduleTimeOfDay() *_scheduleTimeOfDay { _ = "STUB: not implemented"; return nil }

func (u *_scheduleTimeOfDay) String(string string) *_scheduleTimeOfDay {
	_ = "STUB: not implemented"
	return nil
}

func (u *_scheduleTimeOfDay) HourAndMinute(hourandminute types.HourAndMinuteVariant) *_scheduleTimeOfDay {
	_ = "STUB: not implemented"
	return nil
}

func (u *_hourAndMinute) ScheduleTimeOfDayCaster() *types.ScheduleTimeOfDay {
	_ = "STUB: not implemented"
	return nil
}

func (u *_scheduleTimeOfDay) ScheduleTimeOfDayCaster() *types.ScheduleTimeOfDay {
	_ = "STUB: not implemented"
	return nil
}
