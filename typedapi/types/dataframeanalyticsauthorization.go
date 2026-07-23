package types

type DataframeAnalyticsAuthorization struct {
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`

	Roles []string `json:"roles,omitempty"`

	ServiceAccount *string `json:"service_account,omitempty"`
}

func (s *DataframeAnalyticsAuthorization) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsAuthorization() *DataframeAnalyticsAuthorization {
	_ = "STUB: not implemented"
	return nil
}
