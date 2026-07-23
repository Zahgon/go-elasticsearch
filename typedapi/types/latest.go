package types

type Latest struct {
	Sort string `json:"sort"`

	UniqueKey []string `json:"unique_key"`
}

func (s *Latest) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewLatest() *Latest { _ = "STUB: not implemented"; return nil }

type LatestVariant interface {
	LatestCaster() *Latest
}

func (s *Latest) LatestCaster() *Latest { _ = "STUB: not implemented"; return nil }
