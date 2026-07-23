package samllogout

type Request struct {
	RefreshToken *string `json:"refresh_token,omitempty"`

	Token string `json:"token"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
