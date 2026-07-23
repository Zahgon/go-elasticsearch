package types

type AllField struct {
	Analyzer                 string `json:"analyzer"`
	Enabled                  bool   `json:"enabled"`
	OmitNorms                bool   `json:"omit_norms"`
	SearchAnalyzer           string `json:"search_analyzer"`
	Similarity               string `json:"similarity"`
	Store                    bool   `json:"store"`
	StoreTermVectorOffsets   bool   `json:"store_term_vector_offsets"`
	StoreTermVectorPayloads  bool   `json:"store_term_vector_payloads"`
	StoreTermVectorPositions bool   `json:"store_term_vector_positions"`
	StoreTermVectors         bool   `json:"store_term_vectors"`
}

func (s *AllField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAllField() *AllField { _ = "STUB: not implemented"; return nil }

type AllFieldVariant interface {
	AllFieldCaster() *AllField
}

func (s *AllField) AllFieldCaster() *AllField { _ = "STUB: not implemented"; return nil }
