package puttrainedmodeldefinitionpart

type Request struct {
	Definition string `json:"definition"`

	TotalDefinitionLength int64 `json:"total_definition_length"`

	TotalParts int `json:"total_parts"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
