package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uppercaseProcessor struct {
	v *types.UppercaseProcessor
}

func NewUppercaseProcessor() *_uppercaseProcessor { _ = "STUB: not implemented"; return nil }

func (s *_uppercaseProcessor) Field(field string) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) IgnoreMissing(ignoremissing bool) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) TargetField(field string) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) Description(description string) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) If(if_ types.ScriptVariant) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) IgnoreFailure(ignorefailure bool) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) Tag(tag string) *_uppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uppercaseProcessor) UppercaseProcessorCaster() *types.UppercaseProcessor {
	_ = "STUB: not implemented"
	return nil
}
