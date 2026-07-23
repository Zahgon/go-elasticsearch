package types

type TriggerContainer struct {
	Schedule *ScheduleContainer `json:"schedule,omitempty"`
}

func NewTriggerContainer() *TriggerContainer { _ = "STUB: not implemented"; return nil }

type TriggerContainerVariant interface {
	TriggerContainerCaster() *TriggerContainer
}

func (s *TriggerContainer) TriggerContainerCaster() *TriggerContainer {
	_ = "STUB: not implemented"
	return nil
}
