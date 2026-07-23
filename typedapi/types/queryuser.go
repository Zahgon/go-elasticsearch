package types

type QueryUser struct {
	Email      *string      `json:"email,omitempty"`
	Enabled    bool         `json:"enabled"`
	FullName   *string      `json:"full_name,omitempty"`
	Metadata   Metadata     `json:"metadata"`
	ProfileUid *string      `json:"profile_uid,omitempty"`
	Roles      []string     `json:"roles"`
	Sort_      []FieldValue `json:"_sort,omitempty"`
	Username   string       `json:"username"`
}

func (s *QueryUser) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryUser() *QueryUser { _ = "STUB: not implemented"; return nil }
