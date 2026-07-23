package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _openShiftAiTaskSettings struct {
	v *types.OpenShiftAiTaskSettings
}

func NewOpenShiftAiTaskSettings() *_openShiftAiTaskSettings { _ = "STUB: not implemented"; return nil }

func (s *_openShiftAiTaskSettings) ReturnDocuments(returndocuments bool) *_openShiftAiTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiTaskSettings) TopN(topn int) *_openShiftAiTaskSettings {
	_ = "STUB: not implemented"
	return nil
}

func (s *_openShiftAiTaskSettings) OpenShiftAiTaskSettingsCaster() *types.OpenShiftAiTaskSettings {
	_ = "STUB: not implemented"
	return nil
}
