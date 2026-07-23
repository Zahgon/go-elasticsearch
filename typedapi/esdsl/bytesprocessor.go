package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bytesProcessor struct {
	v *types.BytesProcessor
}

func NewBytesProcessor() *_bytesProcessor { _ = "STUB: not implemented"; return nil }

func (s *_bytesProcessor) Field(field string) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) IgnoreMissing(ignoremissing bool) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) TargetField(field string) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) Description(description string) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) If(if_ types.ScriptVariant) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) IgnoreFailure(ignorefailure bool) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_bytesProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) Tag(tag string) *_bytesProcessor { _ = "STUB: not implemented"; return nil }

func (s *_bytesProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_bytesProcessor) BytesProcessorCaster() *types.BytesProcessor {
	_ = "STUB: not implemented"
	return nil
}
