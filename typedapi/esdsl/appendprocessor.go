package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _appendProcessor struct {
	v *types.AppendProcessor
}

func NewAppendProcessor() *_appendProcessor { _ = "STUB: not implemented"; return nil }

func (s *_appendProcessor) AllowDuplicates(allowduplicates bool) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) CopyFrom(field string) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) Field(field string) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) IgnoreEmptyValues(ignoreemptyvalues bool) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) MediaType(mediatype string) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) Value(values ...json.RawMessage) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) Description(description string) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) If(if_ types.ScriptVariant) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) IgnoreFailure(ignorefailure bool) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_appendProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) Tag(tag string) *_appendProcessor { _ = "STUB: not implemented"; return nil }

func (s *_appendProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_appendProcessor) AppendProcessorCaster() *types.AppendProcessor {
	_ = "STUB: not implemented"
	return nil
}
