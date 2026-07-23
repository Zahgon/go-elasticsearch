package types

type ScheduleTriggerEvent struct {
	ScheduledTime DateTime `json:"scheduled_time"`
	TriggeredTime DateTime `json:"triggered_time,omitempty"`
}

func (s *ScheduleTriggerEvent) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewScheduleTriggerEvent() *ScheduleTriggerEvent { _ = "STUB: not implemented"; return nil }

type ScheduleTriggerEventVariant interface {
	ScheduleTriggerEventCaster() *ScheduleTriggerEvent
}

func (s *ScheduleTriggerEvent) ScheduleTriggerEventCaster() *ScheduleTriggerEvent {
	_ = "STUB: not implemented"
	return nil
}
