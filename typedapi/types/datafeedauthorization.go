package types

type DatafeedAuthorization struct {
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`

	Roles []string `json:"roles,omitempty"`

	ServiceAccount *string `json:"service_account,omitempty"`
}

func (s *DatafeedAuthorization) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDatafeedAuthorization() *DatafeedAuthorization { _ = "STUB: not implemented"; return nil }
