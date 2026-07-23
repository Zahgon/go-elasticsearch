package upgradejobsnapshot

type Response struct {
	Completed bool `json:"completed"`

	Node string `json:"node"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
