package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keyValueProcessor struct {
	v *types.KeyValueProcessor
}

func NewKeyValueProcessor(fieldsplit string, valuesplit string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) ExcludeKeys(excludekeys ...string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) Field(field string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) FieldSplit(fieldsplit string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) IgnoreMissing(ignoremissing bool) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) IncludeKeys(includekeys ...string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) Prefix(prefix string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) StripBrackets(stripbrackets bool) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) TargetField(field string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) TrimKey(trimkey string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) TrimValue(trimvalue string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) ValueSplit(valuesplit string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) Description(description string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) If(if_ types.ScriptVariant) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) IgnoreFailure(ignorefailure bool) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) Tag(tag string) *_keyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_keyValueProcessor) KeyValueProcessorCaster() *types.KeyValueProcessor {
	_ = "STUB: not implemented"
	return nil
}
