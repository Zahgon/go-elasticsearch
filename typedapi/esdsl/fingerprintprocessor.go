package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fingerprintdigest"
)

type _fingerprintProcessor struct {
	v *types.FingerprintProcessor
}

func NewFingerprintProcessor() *_fingerprintProcessor { _ = "STUB: not implemented"; return nil }

func (s *_fingerprintProcessor) Fields(fields ...string) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) IgnoreMissing(ignoremissing bool) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) Method(method fingerprintdigest.FingerprintDigest) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) Salt(salt string) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) TargetField(field string) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) Description(description string) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) If(if_ types.ScriptVariant) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) IgnoreFailure(ignorefailure bool) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) OnFailure(onfailures ...types.ProcessorContainerVariant) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) OnFailureValues(onfailurevalues []types.ProcessorContainer) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) Tag(tag string) *_fingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) ProcessorContainerCaster() *types.ProcessorContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fingerprintProcessor) FingerprintProcessorCaster() *types.FingerprintProcessor {
	_ = "STUB: not implemented"
	return nil
}
