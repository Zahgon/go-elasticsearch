package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trimProcessor struct {
	v *types.TrimProcessor
}

func NewTrimProcessor() *_trimProcessor { _ = "STUB: not implemented"; return nil }

func (s *_trimProcessor) Field(field string) *_trimProcessor { _ = "STUB: not implemented"; return nil }

func (s *_trimProcessor) IgnoreMissing(ignoremissing bool) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) TargetField(field string) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) Description(description string) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) If(if_ types.ScriptVariant) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) IgnoreFailure(ignorefailure bool) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_trimProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) Tag(tag string) *_trimProcessor { _ = "STUB: not implemented"; return nil }

func (s *_trimProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trimProcessor) TrimProcessorCaster() *types.TrimProcessor {
	_ = "STUB: not implemented"
	return nil
}
