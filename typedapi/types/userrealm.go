package types

type UserRealm struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *UserRealm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUserRealm() *UserRealm { _ = "STUB: not implemented"; return nil }
