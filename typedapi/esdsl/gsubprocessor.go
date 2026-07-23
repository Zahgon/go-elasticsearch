package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _gsubProcessor struct {
	v *types.GsubProcessor
}

func NewGsubProcessor(pattern string, replacement string) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) Field(field string) *_gsubProcessor { _ = "STUB: not implemented"; return nil }

func (s *_gsubProcessor) IgnoreMissing(ignoremissing bool) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) Pattern(pattern string) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) Replacement(replacement string) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) TargetField(field string) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) Description(description string) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) If(if_ types.ScriptVariant) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) IgnoreFailure(ignorefailure bool) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_gsubProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) Tag(tag string) *_gsubProcessor { _ = "STUB: not implemented"; return nil }

func (s *_gsubProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_gsubProcessor) GsubProcessorCaster() *types.GsubProcessor {
	_ = "STUB: not implemented"
	return nil
}
