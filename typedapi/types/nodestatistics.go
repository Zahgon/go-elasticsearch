package types

type NodeStatistics struct {
	Failed   int          `json:"failed"`
	Failures []ErrorCause `json:"failures,omitempty"`

	Successful int `json:"successful"`

	Total int `json:"total"`
}

func (s *NodeStatistics) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeStatistics() *NodeStatistics { _ = "STUB: not implemented"; return nil }
