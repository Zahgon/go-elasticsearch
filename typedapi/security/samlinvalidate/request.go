package samlinvalidate

type Request struct {
	Acs *string `json:"acs,omitempty"`

	QueryString string `json:"query_string"`

	Realm *string `json:"realm,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
