package types

type UserProfileUser struct {
	Email       *string  `json:"email,omitempty"`
	FullName    *string  `json:"full_name,omitempty"`
	RealmDomain *string  `json:"realm_domain,omitempty"`
	RealmName   string   `json:"realm_name"`
	Roles       []string `json:"roles"`
	Username    string   `json:"username"`
}

func (s *UserProfileUser) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUserProfileUser() *UserProfileUser { _ = "STUB: not implemented"; return nil }
