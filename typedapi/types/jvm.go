package types

type Jvm struct {
	BufferPools map[string]NodeBufferPool `json:"buffer_pools,omitempty"`

	Classes *JvmClasses `json:"classes,omitempty"`

	Gc *GarbageCollector `json:"gc,omitempty"`

	Mem *JvmMemoryStats `json:"mem,omitempty"`

	Threads *JvmThreads `json:"threads,omitempty"`

	Timestamp *int64 `json:"timestamp,omitempty"`

	Uptime *string `json:"uptime,omitempty"`

	UptimeInMillis *int64 `json:"uptime_in_millis,omitempty"`
}

func (s *Jvm) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJvm() *Jvm { _ = "STUB: not implemented"; return nil }
