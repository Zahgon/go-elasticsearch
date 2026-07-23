package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rerouteProcessor struct {
	v *types.RerouteProcessor
}

func NewRerouteProcessor() *_rerouteProcessor { _ = "STUB: not implemented"; return nil }

func (s *_rerouteProcessor) Dataset(datasets ...string) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) Destination(destination string) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) Namespace(namespaces ...string) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) Description(description string) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) If(if_ types.ScriptVariant) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) IgnoreFailure(ignorefailure bool) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) Tag(tag string) *_rerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rerouteProcessor) RerouteProcessorCaster() *types.RerouteProcessor {
	_ = "STUB: not implemented"
	return nil
}
