package clearscroll

type Response struct {
	NumFreed int `json:"num_freed"`

	Succeeded bool `json:"succeeded"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
