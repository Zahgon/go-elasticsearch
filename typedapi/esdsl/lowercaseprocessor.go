package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _lowercaseProcessor struct {
	v *types.LowercaseProcessor
}

func NewLowercaseProcessor() *_lowercaseProcessor { _ = "STUB: not implemented"; return nil }

func (s *_lowercaseProcessor) Field(field string) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) IgnoreMissing(ignoremissing bool) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) TargetField(field string) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) Description(description string) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) If(if_ types.ScriptVariant) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) IgnoreFailure(ignorefailure bool) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) Tag(tag string) *_lowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_lowercaseProcessor) LowercaseProcessorCaster() *types.LowercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}
