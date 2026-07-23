package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _openAITaskSettings struct {
	v *types.OpenAITaskSettings
}

func NewOpenAITaskSettings() *_openAITaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_openAITaskSettings) Headers(headers map[string]string) *_openAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAITaskSettings) AddHeader(key string, value string) *_openAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAITaskSettings) User(user string) *_openAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openAITaskSettings) OpenAITaskSettingsCaster() *types.OpenAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
