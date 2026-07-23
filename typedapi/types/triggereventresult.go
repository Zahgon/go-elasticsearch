package types

type TriggerEventResult struct {
	Manual        TriggerEventContainer `json:"manual"`
	TriggeredTime DateTime              `json:"triggered_time"`
	Type          string                `json:"type"`
}

func (s *TriggerEventResult) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTriggerEventResult() *TriggerEventResult { _ = "STUB: not implemented"; return nil }
