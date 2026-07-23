package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _contextualAITaskSettings struct {
	v *types.ContextualAITaskSettings
}

func NewContextualAITaskSettings() *_contextualAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAITaskSettings) Instruction(instruction string) *_contextualAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAITaskSettings) TopK(topk int) *_contextualAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_contextualAITaskSettings) ContextualAITaskSettingsCaster() *types.ContextualAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
