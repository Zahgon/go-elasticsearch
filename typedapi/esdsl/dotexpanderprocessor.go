package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dotExpanderProcessor struct {
	v *types.DotExpanderProcessor
}

func NewDotExpanderProcessor() *_dotExpanderProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dotExpanderProcessor) Field(field string) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) Override(override bool) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) Path(path string) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) Description(description string) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) If(if_ types.ScriptVariant) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) IgnoreFailure(ignorefailure bool) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) Tag(tag string) *_dotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dotExpanderProcessor) DotExpanderProcessorCaster() *types.DotExpanderProcessor {
	_ = "STUB: not implemented"
	return nil
}
