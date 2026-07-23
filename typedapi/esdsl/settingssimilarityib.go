package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ibdistribution"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/iblambda"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type _settingsSimilarityIb struct {
	v *types.SettingsSimilarityIb
}

func NewSettingsSimilarityIb(distribution ibdistribution.IBDistribution, lambda iblambda.IBLambda, normalization normalization.Normalization) *_settingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityIb) Distribution(distribution ibdistribution.IBDistribution) *_settingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityIb) Lambda(lambda iblambda.IBLambda) *_settingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityIb) Normalization(normalization normalization.Normalization) *_settingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityIb) SettingsSimilarityIbCaster() *types.SettingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}
