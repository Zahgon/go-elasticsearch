package stoptrainedmodeldeployment

type Response struct {
	Stopped bool `json:"stopped"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
