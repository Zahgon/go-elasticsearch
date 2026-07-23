package types

type TransformAuthorization struct {
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`

	Roles []string `json:"roles,omitempty"`

	ServiceAccount *string `json:"service_account,omitempty"`
}

func (s *TransformAuthorization) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransformAuthorization() *TransformAuthorization { _ = "STUB: not implemented"; return nil }
