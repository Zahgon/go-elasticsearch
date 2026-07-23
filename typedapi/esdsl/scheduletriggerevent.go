package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _scheduleTriggerEvent struct {
	v *types.ScheduleTriggerEvent
}

func NewScheduleTriggerEvent() *_scheduleTriggerEvent { _ = "STUB: not implemented"; return nil }

func (s *_scheduleTriggerEvent) ScheduledTime(datetime types.DateTimeVariant) *_scheduleTriggerEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleTriggerEvent) TriggeredTime(datetime types.DateTimeVariant) *_scheduleTriggerEvent {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleTriggerEvent) TriggerEventContainerCaster() *types.TriggerEventContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scheduleTriggerEvent) ScheduleTriggerEventCaster() *types.ScheduleTriggerEvent {
	_ = "STUB: not implemented"
	return nil
}
