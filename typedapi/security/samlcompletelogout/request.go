package samlcompletelogout

type Request struct {
	Content *string `json:"content,omitempty"`

	Ids []string `json:"ids"`

	QueryString *string `json:"query_string,omitempty"`

	Realm string `json:"realm"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
