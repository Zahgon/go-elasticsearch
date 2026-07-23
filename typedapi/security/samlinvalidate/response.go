package samlinvalidate

type Response struct {
	Invalidated int `json:"invalidated"`

	Realm string `json:"realm"`

	Redirect string `json:"redirect"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
