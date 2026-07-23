package oidcprepareauthentication

type Request struct {
	Iss *string `json:"iss,omitempty"`

	LoginHint *string `json:"login_hint,omitempty"`

	Nonce *string `json:"nonce,omitempty"`

	Realm *string `json:"realm,omitempty"`

	State *string `json:"state,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
