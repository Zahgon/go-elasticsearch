package types

type ScriptCache struct {
	CacheEvictions *int64 `json:"cache_evictions,omitempty"`

	CompilationLimitTriggered *int64 `json:"compilation_limit_triggered,omitempty"`

	Compilations *int64  `json:"compilations,omitempty"`
	Context      *string `json:"context,omitempty"`
}

func (s *ScriptCache) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewScriptCache() *ScriptCache { _ = "STUB: not implemented"; return nil }
