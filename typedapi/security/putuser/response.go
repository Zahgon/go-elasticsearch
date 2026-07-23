package putuser

type Response struct {
	Created bool `json:"created"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
