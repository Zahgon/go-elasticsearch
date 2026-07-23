package types

type ClusterProcessCpu struct {
	Percent int `json:"percent"`
}

func (s *ClusterProcessCpu) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterProcessCpu() *ClusterProcessCpu { _ = "STUB: not implemented"; return nil }
