package deletealias

type Response struct {
	Acknowledged bool  `json:"acknowledged"`
	Errors       *bool `json:"errors,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
