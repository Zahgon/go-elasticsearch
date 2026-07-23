package types

type DfsKnnProfile struct {
	Collector             []KnnCollectorResult    `json:"collector"`
	Query                 []KnnQueryProfileResult `json:"query"`
	RewriteTime           int64                   `json:"rewrite_time"`
	VectorOperationsCount *int64                  `json:"vector_operations_count,omitempty"`
}

func (s *DfsKnnProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDfsKnnProfile() *DfsKnnProfile { _ = "STUB: not implemented"; return nil }
