package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/synonymformat"
)

type SynonymTokenFilter struct {
	Expand *bool `json:"expand,omitempty"`

	Format *synonymformat.SynonymFormat `json:"format,omitempty"`

	Lenient *bool `json:"lenient,omitempty"`

	Synonyms []string `json:"synonyms,omitempty"`

	SynonymsPath *string `json:"synonyms_path,omitempty"`

	SynonymsSet []string `json:"synonyms_set,omitempty"`

	Tokenizer *string `json:"tokenizer,omitempty"`
	Type      string  `json:"type,omitempty"`

	Updateable *bool   `json:"updateable,omitempty"`
	Version    *string `json:"version,omitempty"`
}

func (s *SynonymTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SynonymTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSynonymTokenFilter() *SynonymTokenFilter { _ = "STUB: not implemented"; return nil }

type SynonymTokenFilterVariant interface {
	SynonymTokenFilterCaster() *SynonymTokenFilter
}

func (s *SynonymTokenFilter) SynonymTokenFilterCaster() *SynonymTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SynonymTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
