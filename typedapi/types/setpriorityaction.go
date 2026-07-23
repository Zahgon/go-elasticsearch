package types

type SetPriorityAction struct {
	Priority *int `json:"priority,omitempty"`
}

func (s *SetPriorityAction) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSetPriorityAction() *SetPriorityAction { _ = "STUB: not implemented"; return nil }

type SetPriorityActionVariant interface {
	SetPriorityActionCaster() *SetPriorityAction
}

func (s *SetPriorityAction) SetPriorityActionCaster() *SetPriorityAction {
	_ = "STUB: not implemented"
	return nil
}
