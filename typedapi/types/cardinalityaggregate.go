package types

type CardinalityAggregate struct {
	Meta  Metadata `json:"meta,omitempty"`
	Value int64    `json:"value"`
}

func (s *CardinalityAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCardinalityAggregate() *CardinalityAggregate { _ = "STUB: not implemented"; return nil }
