package types

type JvmClasses struct {
	CurrentLoadedCount *int64 `json:"current_loaded_count,omitempty"`

	TotalLoadedCount *int64 `json:"total_loaded_count,omitempty"`

	TotalUnloadedCount *int64 `json:"total_unloaded_count,omitempty"`
}

func (s *JvmClasses) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJvmClasses() *JvmClasses { _ = "STUB: not implemented"; return nil }
