package oidcprepareauthentication

type Response struct {
	Nonce string `json:"nonce"`
	Realm string `json:"realm"`

	Redirect string `json:"redirect"`
	State    string `json:"state"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
