package types

type ReloadDetails struct {
	Index             string   `json:"index"`
	ReloadedAnalyzers []string `json:"reloaded_analyzers"`
	ReloadedNodeIds   []string `json:"reloaded_node_ids"`
}

func (s *ReloadDetails) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReloadDetails() *ReloadDetails { _ = "STUB: not implemented"; return nil }
