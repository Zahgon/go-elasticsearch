package types

type Page struct {
	From *int `json:"from,omitempty"`

	Size *int `json:"size,omitempty"`
}

func (s *Page) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewPage() *Page { _ = "STUB: not implemented"; return nil }

type PageVariant interface {
	PageCaster() *Page
}

func (s *Page) PageCaster() *Page { _ = "STUB: not implemented"; return nil }
