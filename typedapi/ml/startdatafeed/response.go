package startdatafeed

type Response struct {
	Node []string `json:"node"`

	Started bool `json:"started"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
