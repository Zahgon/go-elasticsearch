package stoptrainedmodeldeployment

type Request struct {
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`

	Force *bool `json:"force,omitempty"`

	Id *string `json:"id,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
