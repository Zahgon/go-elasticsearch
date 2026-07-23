package types

type GarbageCollector struct {
	Collectors map[string]GarbageCollectorTotal `json:"collectors,omitempty"`
}

func NewGarbageCollector() *GarbageCollector { _ = "STUB: not implemented"; return nil }
