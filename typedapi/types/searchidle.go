package types

type SearchIdle struct {
	After Duration `json:"after,omitempty"`
}

func (s *SearchIdle) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchIdle() *SearchIdle { _ = "STUB: not implemented"; return nil }

type SearchIdleVariant interface {
	SearchIdleCaster() *SearchIdle
}

func (s *SearchIdle) SearchIdleCaster() *SearchIdle { _ = "STUB: not implemented"; return nil }
