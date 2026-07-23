package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _registeredDomainProcessor struct {
	v *types.RegisteredDomainProcessor
}

func NewRegisteredDomainProcessor() *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) Field(field string) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) IgnoreMissing(ignoremissing bool) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) TargetField(field string) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) Description(description string) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) If(if_ types.ScriptVariant) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) IgnoreFailure(ignorefailure bool) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) Tag(tag string) *_registeredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_registeredDomainProcessor) RegisteredDomainProcessorCaster() *types.RegisteredDomainProcessor {
	_ = "STUB: not implemented"
	return nil
}
