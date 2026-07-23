package types

type SearchableSnapshots struct {
	Available               bool `json:"available"`
	Enabled                 bool `json:"enabled"`
	FullCopyIndicesCount    *int `json:"full_copy_indices_count,omitempty"`
	IndicesCount            int  `json:"indices_count"`
	SharedCacheIndicesCount *int `json:"shared_cache_indices_count,omitempty"`
}

func (s *SearchableSnapshots) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchableSnapshots() *SearchableSnapshots { _ = "STUB: not implemented"; return nil }
