package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfiindependencemeasure"
)

type SettingsSimilarityDfi struct {
	IndependenceMeasure dfiindependencemeasure.DFIIndependenceMeasure `json:"independence_measure"`
	Type                string                                        `json:"type,omitempty"`
}

func (s SettingsSimilarityDfi) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityDfi() *SettingsSimilarityDfi { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityDfiVariant interface {
	SettingsSimilarityDfiCaster() *SettingsSimilarityDfi
}

func (s *SettingsSimilarityDfi) SettingsSimilarityDfiCaster() *SettingsSimilarityDfi {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityDfi) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
