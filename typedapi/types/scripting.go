package types

type Scripting struct {
	CacheEvictions *int64 `json:"cache_evictions,omitempty"`

	CompilationLimitTriggered *int64 `json:"compilation_limit_triggered,omitempty"`

	Compilations *int64 `json:"compilations,omitempty"`

	CompilationsHistory map[string]int64 `json:"compilations_history,omitempty"`
	Contexts            []NodesContext   `json:"contexts,omitempty"`
}

func (s *Scripting) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScripting() *Scripting { _ = "STUB: not implemented"; return nil }
