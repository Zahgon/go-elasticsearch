package types

type SlicedScroll struct {
	Field *string `json:"field,omitempty"`
	Id    string  `json:"id"`
	Max   int     `json:"max"`
}

func (s *SlicedScroll) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSlicedScroll() *SlicedScroll { _ = "STUB: not implemented"; return nil }

type SlicedScrollVariant interface {
	SlicedScrollCaster() *SlicedScroll
}

func (s *SlicedScroll) SlicedScrollCaster() *SlicedScroll { _ = "STUB: not implemented"; return nil }
