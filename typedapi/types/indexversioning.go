package types

type IndexVersioning struct {
	Created       *string `json:"created,omitempty"`
	CreatedString *string `json:"created_string,omitempty"`
}

func (s *IndexVersioning) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexVersioning() *IndexVersioning { _ = "STUB: not implemented"; return nil }

type IndexVersioningVariant interface {
	IndexVersioningCaster() *IndexVersioning
}

func (s *IndexVersioning) IndexVersioningCaster() *IndexVersioning {
	_ = "STUB: not implemented"
	return nil
}
