package types

type NodeInfoIngest struct {
	Processors []NodeInfoIngestProcessor `json:"processors"`
}

func NewNodeInfoIngest() *NodeInfoIngest { _ = "STUB: not implemented"; return nil }
