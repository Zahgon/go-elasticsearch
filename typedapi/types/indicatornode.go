package types

type IndicatorNode struct {
	Name   *string `json:"name,omitempty"`
	NodeId *string `json:"node_id,omitempty"`
}

func (s *IndicatorNode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicatorNode() *IndicatorNode { _ = "STUB: not implemented"; return nil }
