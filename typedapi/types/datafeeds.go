package types

type Datafeeds struct {
	ScrollSize int `json:"scroll_size"`
}

func (s *Datafeeds) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDatafeeds() *Datafeeds { _ = "STUB: not implemented"; return nil }
