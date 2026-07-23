package types

type Http struct {
	Clients []Client `json:"clients,omitempty"`

	CurrentOpen *int `json:"current_open,omitempty"`

	Routes map[string]HttpRoute `json:"routes"`

	TotalOpened *int64 `json:"total_opened,omitempty"`
}

func (s *Http) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewHttp() *Http { _ = "STUB: not implemented"; return nil }
