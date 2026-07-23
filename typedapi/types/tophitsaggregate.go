package types

type TopHitsAggregate struct {
	Hits HitsMetadata `json:"hits"`
	Meta Metadata     `json:"meta,omitempty"`
}

func (s *TopHitsAggregate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTopHitsAggregate() *TopHitsAggregate { _ = "STUB: not implemented"; return nil }
