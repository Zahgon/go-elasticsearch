package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _setProcessor struct {
	v *types.SetProcessor
}

func NewSetProcessor() *_setProcessor { _ = "STUB: not implemented"; return nil }

func (s *_setProcessor) CopyFrom(field string) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) Field(field string) *_setProcessor { _ = "STUB: not implemented"; return nil }

func (s *_setProcessor) IgnoreEmptyValue(ignoreemptyvalue bool) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) MediaType(mediatype string) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) Override(override bool) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) Value(value json.RawMessage) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) Description(description string) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) If(if_ types.ScriptVariant) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) IgnoreFailure(ignorefailure bool) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_setProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) Tag(tag string) *_setProcessor { _ = "STUB: not implemented"; return nil }

func (s *_setProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_setProcessor) SetProcessorCaster() *types.SetProcessor {
	_ = "STUB: not implemented"
	return nil
}
