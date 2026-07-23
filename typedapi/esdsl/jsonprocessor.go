package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jsonprocessorconflictstrategy"
)

type _jsonProcessor struct {
	v *types.JsonProcessor
}

func NewJsonProcessor() *_jsonProcessor { _ = "STUB: not implemented"; return nil }

func (s *_jsonProcessor) AddToRoot(addtoroot bool) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) AddToRootConflictStrategy(addtorootconflictstrategy jsonprocessorconflictstrategy.JsonProcessorConflictStrategy) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) AllowDuplicateKeys(allowduplicatekeys bool) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) Field(field string) *_jsonProcessor { _ = "STUB: not implemented"; return nil }

func (s *_jsonProcessor) TargetField(field string) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) Description(description string) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) If(if_ types.ScriptVariant) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) IgnoreFailure(ignorefailure bool) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_jsonProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) Tag(tag string) *_jsonProcessor { _ = "STUB: not implemented"; return nil }

func (s *_jsonProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_jsonProcessor) JsonProcessorCaster() *types.JsonProcessor {
	_ = "STUB: not implemented"
	return nil
}
