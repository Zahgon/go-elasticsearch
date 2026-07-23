package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dissectProcessor struct {
	v *types.DissectProcessor
}

func NewDissectProcessor(pattern string) *_dissectProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dissectProcessor) AppendSeparator(appendseparator string) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) Field(field string) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) IgnoreMissing(ignoremissing bool) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) Pattern(pattern string) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) Description(description string) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) If(if_ types.ScriptVariant) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) IgnoreFailure(ignorefailure bool) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) Tag(tag string) *_dissectProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dissectProcessor) DissectProcessorCaster() *types.DissectProcessor {
	_ = "STUB: not implemented"
	return nil
}
