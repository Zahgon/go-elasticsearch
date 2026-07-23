package types

type CalendarEvent struct {
	CalendarId *string `json:"calendar_id,omitempty"`

	Description string `json:"description"`

	EndTime DateTime `json:"end_time"`
	EventId *string  `json:"event_id,omitempty"`

	ForceTimeShift *int `json:"force_time_shift,omitempty"`

	SkipModelUpdate *bool `json:"skip_model_update,omitempty"`

	SkipResult *bool `json:"skip_result,omitempty"`

	StartTime DateTime `json:"start_time"`
}

func (s *CalendarEvent) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCalendarEvent() *CalendarEvent { _ = "STUB: not implemented"; return nil }

type CalendarEventVariant interface {
	CalendarEventCaster() *CalendarEvent
}

func (s *CalendarEvent) CalendarEventCaster() *CalendarEvent { _ = "STUB: not implemented"; return nil }
