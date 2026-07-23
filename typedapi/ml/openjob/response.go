package openjob

type Response struct {
	Node   string `json:"node"`
	Opened bool   `json:"opened"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
