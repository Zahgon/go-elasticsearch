package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pipelineProcessor struct {
	v *types.PipelineProcessor
}

func NewPipelineProcessor() *_pipelineProcessor { _ = "STUB: not implemented"; return nil }

func (s *_pipelineProcessor) IgnoreMissingPipeline(ignoremissingpipeline bool) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) Name(name string) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) Description(description string) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) If(if_ types.ScriptVariant) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) IgnoreFailure(ignorefailure bool) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) Tag(tag string) *_pipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_pipelineProcessor) PipelineProcessorCaster() *types.PipelineProcessor {
	_ = "STUB: not implemented"
	return nil
}
