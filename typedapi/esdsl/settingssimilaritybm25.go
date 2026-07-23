package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _settingsSimilarityBm25 struct {
	v *types.SettingsSimilarityBm25
}

func NewSettingsSimilarityBm25() *_settingsSimilarityBm25 { _ = "STUB: not implemented"; return nil }

func (s *_settingsSimilarityBm25) B(b types.Float64) *_settingsSimilarityBm25 {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityBm25) DiscountOverlaps(discountoverlaps bool) *_settingsSimilarityBm25 {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityBm25) K1(k1 types.Float64) *_settingsSimilarityBm25 {
	_ = "STUB: not implemented"
	return nil
}

func (s *_settingsSimilarityBm25) SettingsSimilarityBm25Caster() *types.SettingsSimilarityBm25 {
	_ = "STUB: not implemented"
	return nil
}
