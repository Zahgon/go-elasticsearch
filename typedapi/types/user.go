package types

type User struct {
	Email      *string  `json:"email,omitempty"`
	Enabled    bool     `json:"enabled"`
	FullName   *string  `json:"full_name,omitempty"`
	Metadata   Metadata `json:"metadata"`
	ProfileUid *string  `json:"profile_uid,omitempty"`
	Roles      []string `json:"roles"`
	Username   string   `json:"username"`
}

func (s *User) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewUser() *User { _ = "STUB: not implemented"; return nil }
