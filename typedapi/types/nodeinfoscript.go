package types

type NodeInfoScript struct {
	AllowedTypes               string  `json:"allowed_types"`
	DisableMaxCompilationsRate *string `json:"disable_max_compilations_rate,omitempty"`
}

func (s *NodeInfoScript) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoScript() *NodeInfoScript { _ = "STUB: not implemented"; return nil }
