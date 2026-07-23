package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateProcessor struct {
	v *types.DateProcessor
}

func NewDateProcessor() *_dateProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dateProcessor) Field(field string) *_dateProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dateProcessor) Formats(formats ...string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) Locale(locale string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) OutputFormat(outputformat string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) TargetField(field string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) Timezone(timezone string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) Description(description string) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) If(if_ types.ScriptVariant) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) IgnoreFailure(ignorefailure bool) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_dateProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) Tag(tag string) *_dateProcessor { _ = "STUB: not implemented"; return nil }

func (s *_dateProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateProcessor) DateProcessorCaster() *types.DateProcessor {
	_ = "STUB: not implemented"
	return nil
}
