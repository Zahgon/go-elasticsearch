package updatemodelsnapshot

type Request struct {
	Description *string `json:"description,omitempty"`

	Retain *bool `json:"retain,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
