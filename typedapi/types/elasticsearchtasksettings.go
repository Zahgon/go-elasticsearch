package types

type ElasticsearchTaskSettings struct {
	ReturnDocuments *bool `json:"return_documents,omitempty"`
}

func (s *ElasticsearchTaskSettings) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewElasticsearchTaskSettings() *ElasticsearchTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

type ElasticsearchTaskSettingsVariant interface {
	ElasticsearchTaskSettingsCaster() *ElasticsearchTaskSettings
}

func (s *ElasticsearchTaskSettings) ElasticsearchTaskSettingsCaster() *ElasticsearchTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
