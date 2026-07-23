package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dateIndexNameProcessor struct {
	v *types.DateIndexNameProcessor
}

func NewDateIndexNameProcessor(daterounding string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) DateFormats(dateformats ...string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) DateRounding(daterounding string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) Field(field string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) IndexNameFormat(indexnameformat string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) IndexNamePrefix(indexnameprefix string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) Locale(locale string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) Timezone(timezone string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) Description(description string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) If(if_ types.ScriptVariant) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) IgnoreFailure(ignorefailure bool) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) Tag(tag string) *_dateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dateIndexNameProcessor) DateIndexNameProcessorCaster() *types.DateIndexNameProcessor {
	_ = "STUB: not implemented"
	return nil
}
