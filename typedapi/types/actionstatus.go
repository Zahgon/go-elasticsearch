package types

type ActionStatus struct {
	Ack                     AcknowledgeState `json:"ack"`
	LastExecution           *ExecutionState  `json:"last_execution,omitempty"`
	LastSuccessfulExecution *ExecutionState  `json:"last_successful_execution,omitempty"`
	LastThrottle            *ThrottleState   `json:"last_throttle,omitempty"`
}

func NewActionStatus() *ActionStatus { _ = "STUB: not implemented"; return nil }

type ActionStatusVariant interface {
	ActionStatusCaster() *ActionStatus
}

func (s *ActionStatus) ActionStatusCaster() *ActionStatus { _ = "STUB: not implemented"; return nil }
