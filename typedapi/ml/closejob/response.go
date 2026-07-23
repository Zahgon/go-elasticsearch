package closejob

type Response struct {
	Closed bool `json:"closed"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
