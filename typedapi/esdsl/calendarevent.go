package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _calendarEvent struct {
	v *types.CalendarEvent
}

func NewCalendarEvent(description string) *_calendarEvent { _ = "STUB: not implemented"; return nil }

func (s *_calendarEvent) CalendarId(id string) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) Description(description string) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) EndTime(datetime types.DateTimeVariant) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) EventId(id string) *_calendarEvent { _ = "STUB: not implemented"; return nil }

func (s *_calendarEvent) ForceTimeShift(forcetimeshift int) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) SkipModelUpdate(skipmodelupdate bool) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) SkipResult(skipresult bool) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) StartTime(datetime types.DateTimeVariant) *_calendarEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_calendarEvent) CalendarEventCaster() *types.CalendarEvent {
	_ = "STUB: not implemented"
	return nil
}
