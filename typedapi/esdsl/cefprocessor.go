package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _cefProcessor struct {
	v *types.CefProcessor
}

func NewCefProcessor() *_cefProcessor { _ = "STUB: not implemented"; return nil }

func (s *_cefProcessor) Field(field string) *_cefProcessor { _ = "STUB: not implemented"; return nil }

func (s *_cefProcessor) IgnoreEmptyValues(ignoreemptyvalues bool) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) IgnoreMissing(ignoremissing bool) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) TargetField(field string) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) Timezone(timezone string) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) Description(description string) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) If(if_ types.ScriptVariant) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) IgnoreFailure(ignorefailure bool) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_cefProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) Tag(tag string) *_cefProcessor { _ = "STUB: not implemented"; return nil }

func (s *_cefProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_cefProcessor) CefProcessorCaster() *types.CefProcessor {
	_ = "STUB: not implemented"
	return nil
}
