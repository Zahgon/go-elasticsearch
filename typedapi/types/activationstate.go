package types

type ActivationState struct {
	Active    bool     `json:"active"`
	Timestamp DateTime `json:"timestamp"`
}

func (s *ActivationState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewActivationState() *ActivationState { _ = "STUB: not implemented"; return nil }

type ActivationStateVariant interface {
	ActivationStateCaster() *ActivationState
}

func (s *ActivationState) ActivationStateCaster() *ActivationState {
	_ = "STUB: not implemented"
	return nil
}
