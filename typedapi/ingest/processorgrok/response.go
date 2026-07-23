package processorgrok

type Response struct {
	Patterns map[string]string `json:"patterns"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
