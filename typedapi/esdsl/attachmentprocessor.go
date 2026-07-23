package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _attachmentProcessor struct {
	v *types.AttachmentProcessor
}

func NewAttachmentProcessor() *_attachmentProcessor { _ = "STUB: not implemented"; return nil }

func (s *_attachmentProcessor) Field(field string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) IgnoreMissing(ignoremissing bool) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) IndexedChars(indexedchars int64) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) IndexedCharsField(field string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) MaxFieldBytes(bytesize types.ByteSizeVariant) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) Properties(properties ...string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) RemoveBinary(removebinary bool) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) ResourceName(resourcename string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) TargetField(field string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) Description(description string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) If(if_ types.ScriptVariant) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) IgnoreFailure(ignorefailure bool) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) Tag(tag string) *_attachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_attachmentProcessor) AttachmentProcessorCaster() *types.AttachmentProcessor {
	_ = "STUB: not implemented"
	return nil
}
