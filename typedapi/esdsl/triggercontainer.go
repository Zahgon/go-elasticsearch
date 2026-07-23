package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _triggerContainer struct {
	v *types.TriggerContainer
}

func NewTriggerContainer() *_triggerContainer { _ = "STUB: not implemented"; return nil }

func (s *_triggerContainer) Schedule(schedule types.ScheduleContainerVariant) *_triggerContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_triggerContainer) TriggerContainerCaster() *types.TriggerContainer {
	_ = "STUB: not implemented"
	return nil
}
