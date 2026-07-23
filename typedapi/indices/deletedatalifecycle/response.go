package deletedatalifecycle

type Response struct {
	Acknowledged bool `json:"acknowledged"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
