package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _renameProcessor struct {
	v *types.RenameProcessor
}

func NewRenameProcessor() *_renameProcessor { _ = "STUB: not implemented"; return nil }

func (s *_renameProcessor) Field(field string) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) IgnoreMissing(ignoremissing bool) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) TargetField(field string) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) Description(description string) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) If(if_ types.ScriptVariant) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) IgnoreFailure(ignorefailure bool) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_renameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) Tag(tag string) *_renameProcessor { _ = "STUB: not implemented"; return nil }

func (s *_renameProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_renameProcessor) RenameProcessorCaster() *types.RenameProcessor {
	_ = "STUB: not implemented"
	return nil
}
