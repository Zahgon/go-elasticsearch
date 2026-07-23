package types

type IndicesVersions struct {
	IndexCount        int      `json:"index_count"`
	PrimaryShardCount int      `json:"primary_shard_count"`
	TotalPrimaryBytes int64    `json:"total_primary_bytes"`
	TotalPrimarySize  ByteSize `json:"total_primary_size,omitempty"`
	Version           string   `json:"version"`
}

func (s *IndicesVersions) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndicesVersions() *IndicesVersions { _ = "STUB: not implemented"; return nil }
