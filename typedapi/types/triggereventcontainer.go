package types

type TriggerEventContainer struct {
	Schedule *ScheduleTriggerEvent `json:"schedule,omitempty"`
}

func NewTriggerEventContainer() *TriggerEventContainer { _ = "STUB: not implemented"; return nil }
