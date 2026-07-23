package putfilter

type Request struct {
	Description *string `json:"description,omitempty"`

	Items []string `json:"items,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
