package types

type ClusterIndices struct {
	Analysis *CharFilterTypes `json:"analysis,omitempty"`

	Completion CompletionStats `json:"completion"`

	Count int64 `json:"count"`

	DenseVector DenseVectorStats `json:"dense_vector"`

	Docs DocStats `json:"docs"`

	Fielddata FielddataStats `json:"fielddata"`

	Mappings *FieldTypesMappings `json:"mappings,omitempty"`

	QueryCache QueryCacheStats `json:"query_cache"`

	Search SearchUsageStats `json:"search"`

	Segments SegmentsStats `json:"segments"`

	Shards ClusterIndicesShards `json:"shards"`

	SparseVector SparseVectorStats `json:"sparse_vector"`

	Store StoreStats `json:"store"`

	Versions []IndicesVersions `json:"versions,omitempty"`
}

func (s *ClusterIndices) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClusterIndices() *ClusterIndices { _ = "STUB: not implemented"; return nil }
