package oidcauthenticate

type Response struct {
	AccessToken string `json:"access_token"`

	ExpiresIn int `json:"expires_in"`

	RefreshToken string `json:"refresh_token"`

	Type string `json:"type"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
