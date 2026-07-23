package putbehavioralanalytics

type Response struct {
	Acknowledged bool `json:"acknowledged"`

	Name string `json:"name"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
