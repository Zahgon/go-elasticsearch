package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dailySchedule struct {
	v *types.DailySchedule
}

func NewDailySchedule() *_dailySchedule { _ = "STUB: not implemented"; return nil }

func (s *_dailySchedule) At(ats ...types.ScheduleTimeOfDayVariant) *_dailySchedule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dailySchedule) AtValues(atvalues []types.ScheduleTimeOfDay) *_dailySchedule {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dailySchedule) ScheduleContainerCaster() *types.ScheduleContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dailySchedule) DailyScheduleCaster() *types.DailySchedule {
	_ = "STUB: not implemented"
	return nil
}
