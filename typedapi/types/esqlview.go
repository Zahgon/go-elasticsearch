package types

type ESQLView struct {
	Name string `json:"name"`

	Query string `json:"query"`
}

func (s *ESQLView) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewESQLView() *ESQLView { _ = "STUB: not implemented"; return nil }
