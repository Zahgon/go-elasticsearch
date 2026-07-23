package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _grokProcessor struct {
	v *types.GrokProcessor
}

func NewGrokProcessor() *_grokProcessor { _ = "STUB: not implemented"; return nil }

func (s *_grokProcessor) EcsCompatibility(ecscompatibility string) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) Field(field string) *_grokProcessor { _ = "STUB: not implemented"; return nil }

func (s *_grokProcessor) IgnoreMissing(ignoremissing bool) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) PatternDefinitions(patterndefinitions map[string]string) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) AddPatternDefinition(key string, value string) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) Patterns(patterns ...string) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) TraceMatch(tracematch bool) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) ValidateOnly(validateonly bool) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) Description(description string) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) If(if_ types.ScriptVariant) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) IgnoreFailure(ignorefailure bool) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_grokProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) Tag(tag string) *_grokProcessor { _ = "STUB: not implemented"; return nil }

func (s *_grokProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_grokProcessor) GrokProcessorCaster() *types.GrokProcessor {
	_ = "STUB: not implemented"
	return nil
}
