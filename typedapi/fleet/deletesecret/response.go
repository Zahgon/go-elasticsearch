package deletesecret

type Response struct {
	Deleted bool `json:"deleted"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
