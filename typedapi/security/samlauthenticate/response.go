package samlauthenticate

type Response struct {
	AccessToken string `json:"access_token"`

	ExpiresIn int `json:"expires_in"`

	InResponseTo *string `json:"in_response_to,omitempty"`

	Realm string `json:"realm"`

	RefreshToken string `json:"refresh_token"`

	Username string `json:"username"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
