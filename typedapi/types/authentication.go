package types

type Authentication struct {
	ApiKey              map[string]string   `json:"api_key,omitempty"`
	AuthenticationRealm AuthenticationRealm `json:"authentication_realm"`
	AuthenticationType  string              `json:"authentication_type"`
	Email               *string             `json:"email,omitempty"`
	Enabled             bool                `json:"enabled"`
	FullName            *string             `json:"full_name,omitempty"`
	LookupRealm         AuthenticationRealm `json:"lookup_realm"`
	Metadata            Metadata            `json:"metadata"`
	Roles               []string            `json:"roles"`
	Token               map[string]string   `json:"token,omitempty"`
	Username            string              `json:"username"`
}

func (s *Authentication) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAuthentication() *Authentication { _ = "STUB: not implemented"; return nil }
