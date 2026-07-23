package oidcauthenticate

type Request struct {
	Nonce string `json:"nonce"`

	Realm *string `json:"realm,omitempty"`

	RedirectUri string `json:"redirect_uri"`

	State string `json:"state"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
