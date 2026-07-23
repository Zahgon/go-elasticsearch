package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _csvProcessor struct {
	v *types.CsvProcessor
}

func NewCsvProcessor() *_csvProcessor { _ = "STUB: not implemented"; return nil }

func (s *_csvProcessor) EmptyValue(emptyvalue json.RawMessage) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) Field(field string) *_csvProcessor { _ = "STUB: not implemented"; return nil }

func (s *_csvProcessor) IgnoreMissing(ignoremissing bool) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) Quote(quote string) *_csvProcessor { _ = "STUB: not implemented"; return nil }

func (s *_csvProcessor) Separator(separator string) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) TargetFields(fields ...string) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) Trim(trim bool) *_csvProcessor { _ = "STUB: not implemented"; return nil }

func (s *_csvProcessor) Description(description string) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) If(if_ types.ScriptVariant) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) IgnoreFailure(ignorefailure bool) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_csvProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) Tag(tag string) *_csvProcessor { _ = "STUB: not implemented"; return nil }

func (s *_csvProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_csvProcessor) CsvProcessorCaster() *types.CsvProcessor {
	_ = "STUB: not implemented"
	return nil
}
