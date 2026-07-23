package types

type FetchProfile struct {
	Breakdown   FetchProfileBreakdown `json:"breakdown"`
	Children    []FetchProfile        `json:"children,omitempty"`
	Debug       *FetchProfileDebug    `json:"debug,omitempty"`
	Description string                `json:"description"`
	TimeInNanos int64                 `json:"time_in_nanos"`
	Type        string                `json:"type"`
}

func (s *FetchProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFetchProfile() *FetchProfile { _ = "STUB: not implemented"; return nil }
