package updatename

type Request struct {
	Description *string `json:"description,omitempty"`
	Name        *string `json:"name,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
