package types

type FielddataStats struct {
	Evictions         *int64                      `json:"evictions,omitempty"`
	Fields            map[string]FieldMemoryUsage `json:"fields,omitempty"`
	GlobalOrdinals    GlobalOrdinalsStats         `json:"global_ordinals"`
	MemorySize        ByteSize                    `json:"memory_size,omitempty"`
	MemorySizeInBytes int64                       `json:"memory_size_in_bytes"`
}

func (s *FielddataStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFielddataStats() *FielddataStats { _ = "STUB: not implemented"; return nil }
