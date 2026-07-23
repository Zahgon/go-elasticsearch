package types

type NodeInfoIngestProcessor struct {
	Type string `json:"type"`
}

func (s *NodeInfoIngestProcessor) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoIngestProcessor() *NodeInfoIngestProcessor { _ = "STUB: not implemented"; return nil }
