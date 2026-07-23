package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _foreachProcessor struct {
	v *types.ForeachProcessor
}

func NewForeachProcessor(processor types.ProcessorContainerVariant) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) Field(field string) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) IgnoreMissing(ignoremissing bool) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) Processor(processor types.ProcessorContainerVariant) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) Description(description string) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) If(if_ types.ScriptVariant) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) IgnoreFailure(ignorefailure bool) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) Tag(tag string) *_foreachProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_foreachProcessor) ForeachProcessorCaster() *types.ForeachProcessor {
	_ = "STUB: not implemented"
	return nil
}
