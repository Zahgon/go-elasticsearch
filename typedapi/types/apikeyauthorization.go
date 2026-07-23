package types

type ApiKeyAuthorization struct {
	Id string `json:"id"`

	Name string `json:"name"`
}

func (s *ApiKeyAuthorization) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewApiKeyAuthorization() *ApiKeyAuthorization { _ = "STUB: not implemented"; return nil }
