package types

type SegmentsRecord struct {
	Committed *string `json:"committed,omitempty"`

	Compound *string `json:"compound,omitempty"`

	DocsCount *string `json:"docs.count,omitempty"`

	DocsDeleted *string `json:"docs.deleted,omitempty"`

	Generation *string `json:"generation,omitempty"`

	Id *string `json:"id,omitempty"`

	Index *string `json:"index,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Prirep *string `json:"prirep,omitempty"`

	Searchable *string `json:"searchable,omitempty"`

	Segment *string `json:"segment,omitempty"`

	Shard *string `json:"shard,omitempty"`

	Size ByteSize `json:"size,omitempty"`

	SizeMemory ByteSize `json:"size.memory,omitempty"`

	Version *string `json:"version,omitempty"`
}

func (s *SegmentsRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSegmentsRecord() *SegmentsRecord { _ = "STUB: not implemented"; return nil }
