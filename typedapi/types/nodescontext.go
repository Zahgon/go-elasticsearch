package types

type NodesContext struct {
	CacheEvictions            *int64  `json:"cache_evictions,omitempty"`
	CompilationLimitTriggered *int64  `json:"compilation_limit_triggered,omitempty"`
	Compilations              *int64  `json:"compilations,omitempty"`
	Context                   *string `json:"context,omitempty"`
}

func (s *NodesContext) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodesContext() *NodesContext { _ = "STUB: not implemented"; return nil }
