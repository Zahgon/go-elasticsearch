package types

type NodeInfoJvmMemory struct {
	DirectMax          ByteSize `json:"direct_max,omitempty"`
	DirectMaxInBytes   int64    `json:"direct_max_in_bytes"`
	HeapInit           ByteSize `json:"heap_init,omitempty"`
	HeapInitInBytes    int64    `json:"heap_init_in_bytes"`
	HeapMax            ByteSize `json:"heap_max,omitempty"`
	HeapMaxInBytes     int64    `json:"heap_max_in_bytes"`
	NonHeapInit        ByteSize `json:"non_heap_init,omitempty"`
	NonHeapInitInBytes int64    `json:"non_heap_init_in_bytes"`
	NonHeapMax         ByteSize `json:"non_heap_max,omitempty"`
	NonHeapMaxInBytes  int64    `json:"non_heap_max_in_bytes"`
}

func (s *NodeInfoJvmMemory) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeInfoJvmMemory() *NodeInfoJvmMemory { _ = "STUB: not implemented"; return nil }
