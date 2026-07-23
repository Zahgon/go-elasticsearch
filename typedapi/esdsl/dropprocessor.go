package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dropProcessor struct {
	v *types.DropProcessor
}

func NewDropProcessor() *_dropProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dropProcessor) Description(description string) *_dropProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) If(if_ types.ScriptVariant) *_dropProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) IgnoreFailure(ignorefailure bool) *_dropProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dropProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_dropProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) Tag(tag string) *_dropProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dropProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dropProcessor) DropProcessorCaster() *types.DropProcessor {
	_ = "STUB: not implemented"
	return nil
}
