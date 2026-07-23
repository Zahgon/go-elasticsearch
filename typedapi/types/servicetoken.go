package types

type ServiceToken struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (s *ServiceToken) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewServiceToken() *ServiceToken { _ = "STUB: not implemented"; return nil }
