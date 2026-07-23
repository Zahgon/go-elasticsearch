package types

type NodeBufferPool struct {
	Count *int64 `json:"count,omitempty"`

	TotalCapacity *string `json:"total_capacity,omitempty"`

	TotalCapacityInBytes *int64 `json:"total_capacity_in_bytes,omitempty"`

	Used *string `json:"used,omitempty"`

	UsedInBytes *int64 `json:"used_in_bytes,omitempty"`
}

func (s *NodeBufferPool) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeBufferPool() *NodeBufferPool { _ = "STUB: not implemented"; return nil }
