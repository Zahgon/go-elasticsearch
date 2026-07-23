package types

type AllocationStore struct {
	AllocationId        string `json:"allocation_id"`
	Found               bool   `json:"found"`
	InSync              bool   `json:"in_sync"`
	MatchingSizeInBytes int64  `json:"matching_size_in_bytes"`
	MatchingSyncId      bool   `json:"matching_sync_id"`
	StoreException      string `json:"store_exception"`
}

func (s *AllocationStore) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAllocationStore() *AllocationStore { _ = "STUB: not implemented"; return nil }
