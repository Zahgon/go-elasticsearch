package updatefilter

type Request struct {
	AddItems []string `json:"add_items,omitempty"`

	Description *string `json:"description,omitempty"`

	RemoveItems []string `json:"remove_items,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
