package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/synonymformat"
)

type SynonymGraphTokenFilter struct {
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

func (s *SynonymGraphTokenFilter) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s SynonymGraphTokenFilter) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewSynonymGraphTokenFilter() *SynonymGraphTokenFilter { _ = "STUB: not implemented"; return nil }

type SynonymGraphTokenFilterVariant interface {
	SynonymGraphTokenFilterCaster() *SynonymGraphTokenFilter
}

func (s *SynonymGraphTokenFilter) SynonymGraphTokenFilterCaster() *SynonymGraphTokenFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *SynonymGraphTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	_ = "STUB: not implemented"
	return nil
}
