package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _redactProcessor struct {
	v *types.RedactProcessor
}

func NewRedactProcessor() *_redactProcessor { _ = "STUB: not implemented"; return nil }

func (s *_redactProcessor) Field(field string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) IgnoreMissing(ignoremissing bool) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) PatternDefinitions(patterndefinitions map[string]string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) AddPatternDefinition(key string, value string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) Patterns(patterns ...string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) Prefix(prefix string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) SkipIfUnlicensed(skipifunlicensed bool) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) Suffix(suffix string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) TraceRedact(traceredact bool) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) Description(description string) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) If(if_ types.ScriptVariant) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) IgnoreFailure(ignorefailure bool) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_redactProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) Tag(tag string) *_redactProcessor { _ = "STUB: not implemented"; return nil }

func (s *_redactProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_redactProcessor) RedactProcessorCaster() *types.RedactProcessor {
	_ = "STUB: not implemented"
	return nil
}
