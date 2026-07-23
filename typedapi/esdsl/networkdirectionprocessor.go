package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _networkDirectionProcessor struct {
	v *types.NetworkDirectionProcessor
}

func NewNetworkDirectionProcessor() *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) DestinationIp(field string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) IgnoreMissing(ignoremissing bool) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) InternalNetworks(internalnetworks ...string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) InternalNetworksField(field string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) SourceIp(field string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) TargetField(field string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) Description(description string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) If(if_ types.ScriptVariant) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) IgnoreFailure(ignorefailure bool) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) Tag(tag string) *_networkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_networkDirectionProcessor) NetworkDirectionProcessorCaster() *types.NetworkDirectionProcessor {
	_ = "STUB: not implemented"
	return nil
}
