package post

type Request struct {
	Description *string `json:"description,omitempty"`
	IndexName   *string `json:"index_name,omitempty"`
	IsNative    *bool   `json:"is_native,omitempty"`
	Language    *string `json:"language,omitempty"`
	Name        *string `json:"name,omitempty"`
	ServiceType *string `json:"service_type,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
