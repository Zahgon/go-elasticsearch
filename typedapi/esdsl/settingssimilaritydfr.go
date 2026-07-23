package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfraftereffect"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfrbasicmodel"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type _settingsSimilarityDfr struct {
	v *types.SettingsSimilarityDfr
}

func NewSettingsSimilarityDfr(aftereffect dfraftereffect.DFRAfterEffect, basicmodel dfrbasicmodel.DFRBasicModel, normalization normalization.Normalization) *_settingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfr) AfterEffect(aftereffect dfraftereffect.DFRAfterEffect) *_settingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfr) BasicModel(basicmodel dfrbasicmodel.DFRBasicModel) *_settingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfr) Normalization(normalization normalization.Normalization) *_settingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityDfr) SettingsSimilarityDfrCaster() *types.SettingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}
