package samllogout

type Response struct {
	Redirect string `json:"redirect"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
