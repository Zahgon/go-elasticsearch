package revertmodelsnapshot

type Request struct {
	DeleteInterveningResults *bool `json:"delete_intervening_results,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
