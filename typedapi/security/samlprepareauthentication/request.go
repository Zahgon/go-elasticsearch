package samlprepareauthentication

type Request struct {
	Acs *string `json:"acs,omitempty"`

	Realm *string `json:"realm,omitempty"`

	RelayState *string `json:"relay_state,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
