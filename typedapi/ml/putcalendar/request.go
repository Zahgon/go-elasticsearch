package putcalendar

type Request struct {
	Description *string `json:"description,omitempty"`

	JobIds []string `json:"job_ids,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
