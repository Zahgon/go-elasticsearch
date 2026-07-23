package types

type NodeInfoPath struct {
	Data []string `json:"data,omitempty"`
	Home *string  `json:"home,omitempty"`
	Logs *string  `json:"logs,omitempty"`
	Repo []string `json:"repo,omitempty"`
}

func (s *NodeInfoPath) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoPath() *NodeInfoPath { _ = "STUB: not implemented"; return nil }
