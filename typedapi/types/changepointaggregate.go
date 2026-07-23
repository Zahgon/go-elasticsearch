package types

type ChangePointAggregate struct {
	Bucket *ChangePointBucket `json:"bucket,omitempty"`
	Meta   Metadata           `json:"meta,omitempty"`
	Type   ChangeType         `json:"type"`
}

func (s *ChangePointAggregate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewChangePointAggregate() *ChangePointAggregate { _ = "STUB: not implemented"; return nil }
