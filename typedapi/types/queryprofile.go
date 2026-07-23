package types

type QueryProfile struct {
	Breakdown   QueryBreakdown `json:"breakdown"`
	Children    []QueryProfile `json:"children,omitempty"`
	Description string         `json:"description"`
	TimeInNanos int64          `json:"time_in_nanos"`
	Type        string         `json:"type"`
}

func (s *QueryProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewQueryProfile() *QueryProfile { _ = "STUB: not implemented"; return nil }
