package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _voyageAITaskSettings struct {
	v *types.VoyageAITaskSettings
}

func NewVoyageAITaskSettings() *_voyageAITaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_voyageAITaskSettings) InputType(inputtype string) *_voyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAITaskSettings) ReturnDocuments(returndocuments bool) *_voyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAITaskSettings) TopK(topk int) *_voyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAITaskSettings) Truncation(truncation bool) *_voyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_voyageAITaskSettings) VoyageAITaskSettingsCaster() *types.VoyageAITaskSettings {
	_ = "STUB: not implemented"
	return nil
}
