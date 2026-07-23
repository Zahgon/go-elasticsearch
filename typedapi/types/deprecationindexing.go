package types

type DeprecationIndexing struct {
	Enabled string `json:"enabled"`
}

func (s *DeprecationIndexing) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDeprecationIndexing() *DeprecationIndexing { _ = "STUB: not implemented"; return nil }
