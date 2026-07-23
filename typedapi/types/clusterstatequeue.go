package types

type ClusterStateQueue struct {
	Committed *int64 `json:"committed,omitempty"`

	Pending *int64 `json:"pending,omitempty"`

	Total *int64 `json:"total,omitempty"`
}

func (s *ClusterStateQueue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterStateQueue() *ClusterStateQueue { _ = "STUB: not implemented"; return nil }
