package types

type SecurityRolesDlsBitSetCache struct {
	Count int `json:"count"`

	Evictions int64 `json:"evictions"`

	Hits int64 `json:"hits"`

	HitsTimeInMillis int64 `json:"hits_time_in_millis"`

	Memory ByteSize `json:"memory,omitempty"`

	MemoryInBytes uint64 `json:"memory_in_bytes"`

	Misses int64 `json:"misses"`

	MissesTimeInMillis int64 `json:"misses_time_in_millis"`
}

func (s *SecurityRolesDlsBitSetCache) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSecurityRolesDlsBitSetCache() *SecurityRolesDlsBitSetCache {
	_ = "STUB: not implemented"
	return nil
}
