package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/scriptlanguage"
)

type _scriptProcessor struct {
	v *types.ScriptProcessor
}

func NewScriptProcessor() *_scriptProcessor { _ = "STUB: not implemented"; return nil }

func (s *_scriptProcessor) Id(id string) *_scriptProcessor { _ = "STUB: not implemented"; return nil }

func (s *_scriptProcessor) Lang(lang scriptlanguage.ScriptLanguage) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) Params(params map[string]json.RawMessage) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) AddParam(key string, value json.RawMessage) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) Source(scriptsource types.ScriptSourceVariant) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) Description(description string) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) If(if_ types.ScriptVariant) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) IgnoreFailure(ignorefailure bool) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_scriptProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) Tag(tag string) *_scriptProcessor { _ = "STUB: not implemented"; return nil }

func (s *_scriptProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_scriptProcessor) ScriptProcessorCaster() *types.ScriptProcessor {
	_ = "STUB: not implemented"
	return nil
}
