package types

type SegmentsStats struct {
	Count int `json:"count"`

	DocValuesMemory ByteSize `json:"doc_values_memory,omitempty"`

	DocValuesMemoryInBytes int64 `json:"doc_values_memory_in_bytes"`

	FileSizes map[string]ShardFileSizeInfo `json:"file_sizes"`

	FixedBitSet ByteSize `json:"fixed_bit_set,omitempty"`

	FixedBitSetMemoryInBytes int64 `json:"fixed_bit_set_memory_in_bytes"`

	IndexWriterMemory ByteSize `json:"index_writer_memory,omitempty"`

	IndexWriterMemoryInBytes int64 `json:"index_writer_memory_in_bytes"`

	MaxUnsafeAutoIdTimestamp int64 `json:"max_unsafe_auto_id_timestamp"`

	Memory ByteSize `json:"memory,omitempty"`

	MemoryInBytes int64 `json:"memory_in_bytes"`

	NormsMemory ByteSize `json:"norms_memory,omitempty"`

	NormsMemoryInBytes int64 `json:"norms_memory_in_bytes"`

	PointsMemory ByteSize `json:"points_memory,omitempty"`

	PointsMemoryInBytes int64 `json:"points_memory_in_bytes"`

	StoredFieldsMemory ByteSize `json:"stored_fields_memory,omitempty"`

	StoredFieldsMemoryInBytes int64 `json:"stored_fields_memory_in_bytes"`

	TermVectorsMemory ByteSize `json:"term_vectors_memory,omitempty"`

	TermVectorsMemoryInBytes int64 `json:"term_vectors_memory_in_bytes"`

	TermsMemory ByteSize `json:"terms_memory,omitempty"`

	TermsMemoryInBytes int64 `json:"terms_memory_in_bytes"`

	VersionMapMemory ByteSize `json:"version_map_memory,omitempty"`

	VersionMapMemoryInBytes int64 `json:"version_map_memory_in_bytes"`
}

func (s *SegmentsStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSegmentsStats() *SegmentsStats { _ = "STUB: not implemented"; return nil }
