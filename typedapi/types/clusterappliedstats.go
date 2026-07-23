package types

type ClusterAppliedStats struct {
	Recordings []Recording `json:"recordings,omitempty"`
}

func NewClusterAppliedStats() *ClusterAppliedStats { _ = "STUB: not implemented"; return nil }
