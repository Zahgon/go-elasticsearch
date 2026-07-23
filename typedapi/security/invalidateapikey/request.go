package invalidateapikey

type Request struct {
	Id *string `json:"id,omitempty"`

	Ids []string `json:"ids,omitempty"`

	Name *string `json:"name,omitempty"`

	Owner *bool `json:"owner,omitempty"`

	RealmName *string `json:"realm_name,omitempty"`

	Username *string `json:"username,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
