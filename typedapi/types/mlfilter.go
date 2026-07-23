package types

type MLFilter struct {
	Description *string `json:"description,omitempty"`

	FilterId string `json:"filter_id"`

	Items []string `json:"items"`
}

func (s *MLFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMLFilter() *MLFilter { _ = "STUB: not implemented"; return nil }
