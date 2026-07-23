package startjob

type Response struct {
	Started bool `json:"started"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
