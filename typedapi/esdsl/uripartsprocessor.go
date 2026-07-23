package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _uriPartsProcessor struct {
	v *types.UriPartsProcessor
}

func NewUriPartsProcessor() *_uriPartsProcessor { _ = "STUB: not implemented"; return nil }

func (s *_uriPartsProcessor) Field(field string) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) IgnoreMissing(ignoremissing bool) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) KeepOriginal(keeporiginal bool) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) RemoveIfSuccessful(removeifsuccessful bool) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) TargetField(field string) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) Description(description string) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) If(if_ types.ScriptVariant) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) IgnoreFailure(ignorefailure bool) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) Tag(tag string) *_uriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_uriPartsProcessor) UriPartsProcessorCaster() *types.UriPartsProcessor {
	_ = "STUB: not implemented"
	return nil
}
