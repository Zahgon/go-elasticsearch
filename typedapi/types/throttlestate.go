package types

type ThrottleState struct {
	Reason    string   `json:"reason"`
	Timestamp DateTime `json:"timestamp"`
}

func (s *ThrottleState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewThrottleState() *ThrottleState { _ = "STUB: not implemented"; return nil }

type ThrottleStateVariant interface {
	ThrottleStateCaster() *ThrottleState
}

func (s *ThrottleState) ThrottleStateCaster() *ThrottleState { _ = "STUB: not implemented"; return nil }
