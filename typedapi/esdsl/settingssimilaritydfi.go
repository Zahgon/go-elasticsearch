package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfiindependencemeasure"
)

type _settingsSimilarityDfi struct {
	v *types.SettingsSimilarityDfi
}

func NewSettingsSimilarityDfi(independencemeasure dfiindependencemeasure.DFIIndependenceMeasure) *_settingsSimilarityDfi {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfi) IndependenceMeasure(independencemeasure dfiindependencemeasure.DFIIndependenceMeasure) *_settingsSimilarityDfi {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfi) SettingsSimilarityDfiCaster() *types.SettingsSimilarityDfi {
	_ = "STUB: not implemented"
	return nil
}
