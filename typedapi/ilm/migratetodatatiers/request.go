package migratetodatatiers

type Request struct {
	LegacyTemplateToDelete *string `json:"legacy_template_to_delete,omitempty"`
	NodeAttribute          *string `json:"node_attribute,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
