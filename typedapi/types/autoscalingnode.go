package types

type AutoscalingNode struct {
	Name string `json:"name"`
}

func (s *AutoscalingNode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAutoscalingNode() *AutoscalingNode { _ = "STUB: not implemented"; return nil }
