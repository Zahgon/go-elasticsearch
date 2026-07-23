package allocationexplain

type Request struct {
	CurrentNode *string `json:"current_node,omitempty"`

	Index *string `json:"index,omitempty"`

	Primary *bool `json:"primary,omitempty"`

	Shard *int `json:"shard,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
