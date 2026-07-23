package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _hourlySchedule struct {
	v *types.HourlySchedule
}

func NewHourlySchedule() *_hourlySchedule { _ = "STUB: not implemented"; return nil }

func (s *_hourlySchedule) Minute(minutes ...int) *_hourlySchedule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hourlySchedule) ScheduleContainerCaster() *types.ScheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_hourlySchedule) HourlyScheduleCaster() *types.HourlySchedule {
	_ = "STUB: not implemented"
	return nil
}
