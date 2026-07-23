package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/useragentproperty"
)

type _userAgentProcessor struct {
	v *types.UserAgentProcessor
}

func NewUserAgentProcessor() *_userAgentProcessor { _ = "STUB: not implemented"; return nil }

func (s *_userAgentProcessor) ExtractDeviceType(extractdevicetype bool) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) Field(field string) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) IgnoreMissing(ignoremissing bool) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) Properties(properties ...useragentproperty.UserAgentProperty) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) RegexFile(regexfile string) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) TargetField(field string) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) Description(description string) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) If(if_ types.ScriptVariant) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) IgnoreFailure(ignorefailure bool) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) Tag(tag string) *_userAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_userAgentProcessor) UserAgentProcessorCaster() *types.UserAgentProcessor {
	_ = "STUB: not implemented"
	return nil
}
