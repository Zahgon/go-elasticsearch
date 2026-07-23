package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _removeProcessor struct {
	v *types.RemoveProcessor
}

func NewRemoveProcessor() *_removeProcessor { _ = "STUB: not implemented"; return nil }

func (s *_removeProcessor) Field(fields ...string) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) IgnoreMissing(ignoremissing bool) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) Keep(fields ...string) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) Description(description string) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) If(if_ types.ScriptVariant) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) IgnoreFailure(ignorefailure bool) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_removeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) Tag(tag string) *_removeProcessor { _ = "STUB: not implemented"; return nil }

func (s *_removeProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_removeProcessor) RemoveProcessorCaster() *types.RemoveProcessor {
	_ = "STUB: not implemented"
	return nil
}
