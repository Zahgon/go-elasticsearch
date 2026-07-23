package getsecret

type Response struct {
	Id    string `json:"id"`
	Value string `json:"value"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
