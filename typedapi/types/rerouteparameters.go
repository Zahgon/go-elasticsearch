package types

type RerouteParameters struct {
	AllowPrimary bool    `json:"allow_primary"`
	FromNode     *string `json:"from_node,omitempty"`
	Index        string  `json:"index"`
	Node         string  `json:"node"`
	Shard        int     `json:"shard"`
	ToNode       *string `json:"to_node,omitempty"`
}

func (s *RerouteParameters) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRerouteParameters() *RerouteParameters { _ = "STUB: not implemented"; return nil }
