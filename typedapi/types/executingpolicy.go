package types

type ExecutingPolicy struct {
	Name string   `json:"name"`
	Task TaskInfo `json:"task"`
}

func (s *ExecutingPolicy) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExecutingPolicy() *ExecutingPolicy { _ = "STUB: not implemented"; return nil }
