package updateapikeyid

type Request struct {
	ApiKeyId       *string `json:"api_key_id,omitempty"`
	ApiKeySecretId *string `json:"api_key_secret_id,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
