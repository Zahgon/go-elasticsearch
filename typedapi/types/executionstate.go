package types

type ExecutionState struct {
	Reason     *string  `json:"reason,omitempty"`
	Successful bool     `json:"successful"`
	Timestamp  DateTime `json:"timestamp"`
}

func (s *ExecutionState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExecutionState() *ExecutionState { _ = "STUB: not implemented"; return nil }

type ExecutionStateVariant interface {
	ExecutionStateCaster() *ExecutionState
}

func (s *ExecutionState) ExecutionStateCaster() *ExecutionState {
	_ = "STUB: not implemented"
	return nil
}
