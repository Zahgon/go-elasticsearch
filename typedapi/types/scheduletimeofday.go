package types

type ScheduleTimeOfDay any

type ScheduleTimeOfDayVariant interface {
	ScheduleTimeOfDayCaster() *ScheduleTimeOfDay
}
