package updatesettings

type Request struct {
	IndexAutoExpandReplicas *string `json:"index.auto_expand_replicas,omitempty"`
	IndexNumberOfReplicas   *int    `json:"index.number_of_replicas,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
