package types

type NodeJvmInfo struct {
	GcCollectors                          []string          `json:"gc_collectors"`
	InputArguments                        []string          `json:"input_arguments"`
	Mem                                   NodeInfoJvmMemory `json:"mem"`
	MemoryPools                           []string          `json:"memory_pools"`
	Pid                                   int               `json:"pid"`
	StartTimeInMillis                     int64             `json:"start_time_in_millis"`
	UsingBundledJdk                       bool              `json:"using_bundled_jdk"`
	UsingCompressedOrdinaryObjectPointers *string           `json:"using_compressed_ordinary_object_pointers,omitempty"`
	Version                               string            `json:"version"`
	VmName                                string            `json:"vm_name"`
	VmVendor                              string            `json:"vm_vendor"`
	VmVersion                             string            `json:"vm_version"`
}

func (s *NodeJvmInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeJvmInfo() *NodeJvmInfo { _ = "STUB: not implemented"; return nil }
