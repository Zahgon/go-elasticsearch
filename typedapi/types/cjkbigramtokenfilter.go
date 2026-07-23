package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cjkbigramignoredscript"
)

type CjkBigramTokenFilter struct {
	IgnoredScripts []cjkbigramignoredscript.CjkBigramIgnoredScript `json:"ignored_scripts,omitempty"`

	OutputUnigrams *bool   `json:"output_unigrams,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *CjkBigramTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s CjkBigramTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewCjkBigramTokenFilter() *CjkBigramTokenFilter { _ = "STUB: not implemented"; return nil }

type CjkBigramTokenFilterVariant interface {
	CjkBigramTokenFilterCaster() *CjkBigramTokenFilter
}

func (s *CjkBigramTokenFilter) CjkBigramTokenFilterCaster() *CjkBigramTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *CjkBigramTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
