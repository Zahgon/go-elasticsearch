package types

type TermVectorsFilter struct {
	MaxDocFreq *int `json:"max_doc_freq,omitempty"`

	MaxNumTerms *int `json:"max_num_terms,omitempty"`

	MaxTermFreq *int `json:"max_term_freq,omitempty"`

	MaxWordLength *int `json:"max_word_length,omitempty"`

	MinDocFreq *int `json:"min_doc_freq,omitempty"`

	MinTermFreq *int `json:"min_term_freq,omitempty"`

	MinWordLength *int `json:"min_word_length,omitempty"`
}

func (s *TermVectorsFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermVectorsFilter() *TermVectorsFilter { _ = "STUB: not implemented"; return nil }

type TermVectorsFilterVariant interface {
	TermVectorsFilterCaster() *TermVectorsFilter
}

func (s *TermVectorsFilter) TermVectorsFilterCaster() *TermVectorsFilter {
	_ = "STUB: not implemented"
	return nil
}
