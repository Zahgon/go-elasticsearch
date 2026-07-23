package types

type AuthenticationRealm struct {
	Domain *string `json:"domain,omitempty"`
	Name   string  `json:"name"`
	Type   string  `json:"type"`
}

func (s *AuthenticationRealm) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAuthenticationRealm() *AuthenticationRealm { _ = "STUB: not implemented"; return nil }
