package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureOpenAITaskSettings struct {
	v *types.AzureOpenAITaskSettings
}

func NewAzureOpenAITaskSettings() *_azureOpenAITaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_azureOpenAITaskSettings) Headers(headers map[string]string) *_azureOpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAITaskSettings) AddHeader(key string, value string) *_azureOpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAITaskSettings) User(user string) *_azureOpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_azureOpenAITaskSettings) AzureOpenAITaskSettingsCaster() *types.AzureOpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
