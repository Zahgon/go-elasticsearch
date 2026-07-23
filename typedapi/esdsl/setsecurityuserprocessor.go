package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _setSecurityUserProcessor struct {
	v *types.SetSecurityUserProcessor
}

func NewSetSecurityUserProcessor() *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) Field(field string) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) Properties(properties ...string) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) Description(description string) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) If(if_ types.ScriptVariant) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) IgnoreFailure(ignorefailure bool) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) Tag(tag string) *_setSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setSecurityUserProcessor) SetSecurityUserProcessorCaster() *types.SetSecurityUserProcessor {
	_ = "STUB: not implemented"
	return nil
}
