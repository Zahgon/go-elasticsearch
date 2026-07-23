package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/customserviceinputtype"
)

type CustomServiceSettings struct {
	BatchSize *int `json:"batch_size,omitempty"`

	Headers map[string]string `json:"headers,omitempty"`

	InputType map[customserviceinputtype.CustomServiceInputType]string `json:"input_type,omitempty"`

	QueryParameters [][]string `json:"query_parameters,omitempty"`

	Request CustomRequestParams `json:"request"`

	Response CustomResponseParams `json:"response"`

	SecretParameters map[string]string `json:"secret_parameters"`

	Url *string `json:"url,omitempty"`
}

func (s *CustomServiceSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCustomServiceSettings() *CustomServiceSettings { _ = "STUB: not implemented"; return nil }

type CustomServiceSettingsVariant interface {
	CustomServiceSettingsCaster() *CustomServiceSettings
}

func (s *CustomServiceSettings) CustomServiceSettingsCaster() *CustomServiceSettings {
	_ = "STUB: not implemented"
	return nil
}
