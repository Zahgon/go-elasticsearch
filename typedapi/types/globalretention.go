package types

type GlobalRetention struct {
	DefaultRetention Duration `json:"default_retention,omitempty"`
	MaxRetention     Duration `json:"max_retention,omitempty"`
}

func (s *GlobalRetention) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewGlobalRetention() *GlobalRetention { _ = "STUB: not implemented"; return nil }
