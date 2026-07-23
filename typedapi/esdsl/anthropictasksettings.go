package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _anthropicTaskSettings struct {
	v *types.AnthropicTaskSettings
}

func NewAnthropicTaskSettings(maxtokens int) *_anthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicTaskSettings) MaxTokens(maxtokens int) *_anthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicTaskSettings) Temperature(temperature float32) *_anthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicTaskSettings) TopK(topk int) *_anthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicTaskSettings) TopP(topp float32) *_anthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_anthropicTaskSettings) AnthropicTaskSettingsCaster() *types.AnthropicTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
