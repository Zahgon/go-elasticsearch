package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
)

type _sortProcessor struct {
	v *types.SortProcessor
}

func NewSortProcessor() *_sortProcessor { _ = "STUB: not implemented"; return nil }

func (s *_sortProcessor) Field(field string) *_sortProcessor { _ = "STUB: not implemented"; return nil }

func (s *_sortProcessor) Order(order sortorder.SortOrder) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) TargetField(field string) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) Description(description string) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) If(if_ types.ScriptVariant) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) IgnoreFailure(ignorefailure bool) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_sortProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) Tag(tag string) *_sortProcessor { _ = "STUB: not implemented"; return nil }

func (s *_sortProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sortProcessor) SortProcessorCaster() *types.SortProcessor {
	_ = "STUB: not implemented"
	return nil
}
