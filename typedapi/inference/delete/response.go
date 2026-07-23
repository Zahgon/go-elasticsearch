package delete

type Response struct {
	Acknowledged bool     `json:"acknowledged"`
	Pipelines    []string `json:"pipelines"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
