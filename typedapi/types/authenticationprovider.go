package types

type AuthenticationProvider struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

func (s *AuthenticationProvider) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAuthenticationProvider() *AuthenticationProvider { _ = "STUB: not implemented"; return nil }
