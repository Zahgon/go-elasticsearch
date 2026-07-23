package samlprepareauthentication

type Response struct {
	Id string `json:"id"`

	Realm string `json:"realm"`

	Redirect string `json:"redirect"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
