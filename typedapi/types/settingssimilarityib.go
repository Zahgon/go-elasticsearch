package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ibdistribution"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/iblambda"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type SettingsSimilarityIb struct {
	Distribution  ibdistribution.IBDistribution `json:"distribution"`
	Lambda        iblambda.IBLambda             `json:"lambda"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

func (s SettingsSimilarityIb) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityIb() *SettingsSimilarityIb { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityIbVariant interface {
	SettingsSimilarityIbCaster() *SettingsSimilarityIb
}

func (s *SettingsSimilarityIb) SettingsSimilarityIbCaster() *SettingsSimilarityIb {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityIb) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
