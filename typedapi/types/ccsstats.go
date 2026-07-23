package types

type CCSStats struct {
	Clusters map[string]RemoteClusterInfo `json:"clusters,omitempty"`

	Esql_ *CCSUsageStats `json:"_esql,omitempty"`

	Search_ CCSUsageStats `json:"_search"`
}

func NewCCSStats() *CCSStats { _ = "STUB: not implemented"; return nil }
