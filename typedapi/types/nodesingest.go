package types

type NodesIngest struct {
	Pipelines map[string]IngestStats `json:"pipelines,omitempty"`

	Total *IngestTotal `json:"total,omitempty"`
}

func NewNodesIngest() *NodesIngest { _ = "STUB: not implemented"; return nil }
