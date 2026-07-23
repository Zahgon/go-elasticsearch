package types

type AuthenticatedUser struct {
	AuthenticationProvider *AuthenticationProvider `json:"authentication_provider,omitempty"`
	AuthenticationRealm    UserRealm               `json:"authentication_realm"`
	AuthenticationType     string                  `json:"authentication_type"`
	Email                  *string                 `json:"email,omitempty"`
	Enabled                bool                    `json:"enabled"`
	FullName               *string                 `json:"full_name,omitempty"`
	LookupRealm            UserRealm               `json:"lookup_realm"`
	Metadata               Metadata                `json:"metadata"`
	ProfileUid             *string                 `json:"profile_uid,omitempty"`
	Roles                  []string                `json:"roles"`
	Username               string                  `json:"username"`
}

func (s *AuthenticatedUser) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAuthenticatedUser() *AuthenticatedUser { _ = "STUB: not implemented"; return nil }
