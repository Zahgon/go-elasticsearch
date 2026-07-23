package types

type IndexField struct {
	Enabled bool `json:"enabled"`
}

func (s *IndexField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexField() *IndexField { _ = "STUB: not implemented"; return nil }

type IndexFieldVariant interface {
	IndexFieldCaster() *IndexField
}

func (s *IndexField) IndexFieldCaster() *IndexField { _ = "STUB: not implemented"; return nil }
