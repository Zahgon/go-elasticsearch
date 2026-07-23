package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _urlDecodeProcessor struct {
	v *types.UrlDecodeProcessor
}

func NewUrlDecodeProcessor() *_urlDecodeProcessor { _ = "STUB: not implemented"; return nil }

func (s *_urlDecodeProcessor) Field(field string) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) IgnoreMissing(ignoremissing bool) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) TargetField(field string) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) Description(description string) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) If(if_ types.ScriptVariant) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) IgnoreFailure(ignorefailure bool) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) Tag(tag string) *_urlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_urlDecodeProcessor) UrlDecodeProcessorCaster() *types.UrlDecodeProcessor {
	_ = "STUB: not implemented"
	return nil
}
