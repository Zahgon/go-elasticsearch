package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _terminateProcessor struct {
	v *types.TerminateProcessor
}

func NewTerminateProcessor() *_terminateProcessor { _ = "STUB: not implemented"; return nil }

func (s *_terminateProcessor) Description(description string) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) If(if_ types.ScriptVariant) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) IgnoreFailure(ignorefailure bool) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) Tag(tag string) *_terminateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_terminateProcessor) TerminateProcessorCaster() *types.TerminateProcessor {
	_ = "STUB: not implemented"
	return nil
}
