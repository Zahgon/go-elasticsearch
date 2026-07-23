package getsynonymrule

type Response struct {
	Id string `json:"id"`

	Synonyms string `json:"synonyms"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
