package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _splitProcessor struct {
	v *types.SplitProcessor
}

func NewSplitProcessor(separator string) *_splitProcessor { _ = "STUB: not implemented"; return nil }

func (s *_splitProcessor) Field(field string) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) IgnoreMissing(ignoremissing bool) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) PreserveTrailing(preservetrailing bool) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) Separator(separator string) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) TargetField(field string) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) Description(description string) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) If(if_ types.ScriptVariant) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) IgnoreFailure(ignorefailure bool) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_splitProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) Tag(tag string) *_splitProcessor { _ = "STUB: not implemented"; return nil }

func (s *_splitProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_splitProcessor) SplitProcessorCaster() *types.SplitProcessor {
	_ = "STUB: not implemented"
	return nil
}
