package invalidatetoken

type Request struct {
	RealmName *string `json:"realm_name,omitempty"`

	RefreshToken *string `json:"refresh_token,omitempty"`

	Token *string `json:"token,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
