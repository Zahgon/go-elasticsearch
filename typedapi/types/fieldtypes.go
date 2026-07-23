package types

type FieldTypes struct {
	Count int `json:"count"`

	IndexCount int `json:"index_count"`

	IndexedVectorCount *int `json:"indexed_vector_count,omitempty"`

	IndexedVectorDimMax *int `json:"indexed_vector_dim_max,omitempty"`

	IndexedVectorDimMin *int `json:"indexed_vector_dim_min,omitempty"`

	Name string `json:"name"`

	ScriptCount *int `json:"script_count,omitempty"`

	VectorElementTypeCount map[string]int `json:"vector_element_type_count,omitempty"`

	VectorIndexTypeCount map[string]int `json:"vector_index_type_count,omitempty"`

	VectorSimilarityTypeCount map[string]int `json:"vector_similarity_type_count,omitempty"`
}

func (s *FieldTypes) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFieldTypes() *FieldTypes { _ = "STUB: not implemented"; return nil }
