package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _joinProcessor struct {
	v *types.JoinProcessor
}

func NewJoinProcessor(separator string) *_joinProcessor { _ = "STUB: not implemented"; return nil }

func (s *_joinProcessor) Field(field string) *_joinProcessor { _ = "STUB: not implemented"; return nil }

func (s *_joinProcessor) Separator(separator string) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) TargetField(field string) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) Description(description string) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) If(if_ types.ScriptVariant) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) IgnoreFailure(ignorefailure bool) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_joinProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) Tag(tag string) *_joinProcessor { _ = "STUB: not implemented"; return nil }

func (s *_joinProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_joinProcessor) JoinProcessorCaster() *types.JoinProcessor {
	_ = "STUB: not implemented"
	return nil
}
