package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _huggingFaceTaskSettings struct {
	v *types.HuggingFaceTaskSettings
}

func NewHuggingFaceTaskSettings() *_huggingFaceTaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_huggingFaceTaskSettings) ReturnDocuments(returndocuments bool) *_huggingFaceTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceTaskSettings) TopN(topn int) *_huggingFaceTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_huggingFaceTaskSettings) HuggingFaceTaskSettingsCaster() *types.HuggingFaceTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
