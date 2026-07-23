package startdataframeanalytics

type Response struct {
	Acknowledged bool `json:"acknowledged"`

	Node string `json:"node"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
