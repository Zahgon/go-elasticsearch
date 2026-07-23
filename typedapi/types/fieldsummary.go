package types

type FieldSummary struct {
	Any           uint          `json:"any"`
	DocValues     uint          `json:"doc_values"`
	InvertedIndex InvertedIndex `json:"inverted_index"`
	KnnVectors    uint          `json:"knn_vectors"`
	Norms         uint          `json:"norms"`
	Points        uint          `json:"points"`
	StoredFields  uint          `json:"stored_fields"`
	TermVectors   uint          `json:"term_vectors"`
}

func NewFieldSummary() *FieldSummary { _ = "STUB: not implemented"; return nil }
