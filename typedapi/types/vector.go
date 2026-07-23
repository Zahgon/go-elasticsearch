package types

type Vector struct {
	Available               bool `json:"available"`
	DenseVectorDimsAvgCount int  `json:"dense_vector_dims_avg_count"`
	DenseVectorFieldsCount  int  `json:"dense_vector_fields_count"`
	Enabled                 bool `json:"enabled"`
	SparseVectorFieldsCount *int `json:"sparse_vector_fields_count,omitempty"`
}

func (s *Vector) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewVector() *Vector { _ = "STUB: not implemented"; return nil }
