package types

type ClusterNode struct {
	Name string `json:"name"`
}

func (s *ClusterNode) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterNode() *ClusterNode { _ = "STUB: not implemented"; return nil }
