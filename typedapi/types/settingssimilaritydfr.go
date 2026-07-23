package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfraftereffect"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dfrbasicmodel"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/normalization"
)

type SettingsSimilarityDfr struct {
	AfterEffect   dfraftereffect.DFRAfterEffect `json:"after_effect"`
	BasicModel    dfrbasicmodel.DFRBasicModel   `json:"basic_model"`
	Normalization normalization.Normalization   `json:"normalization"`
	Type          string                        `json:"type,omitempty"`
}

func (s SettingsSimilarityDfr) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSettingsSimilarityDfr() *SettingsSimilarityDfr { _ = "STUB: not implemented"; return nil }

type SettingsSimilarityDfrVariant interface {
	SettingsSimilarityDfrCaster() *SettingsSimilarityDfr
}

func (s *SettingsSimilarityDfr) SettingsSimilarityDfrCaster() *SettingsSimilarityDfr {
	_ = "STUB: not implemented"
	return nil
}

func (s *SettingsSimilarityDfr) SettingsSimilarityCaster() *SettingsSimilarity {
	_ = "STUB: not implemented"
	return nil
}
