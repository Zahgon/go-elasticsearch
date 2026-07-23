package types

type SearchProfile struct {
	Collector   []Collector    `json:"collector"`
	Query       []QueryProfile `json:"query"`
	RewriteTime int64          `json:"rewrite_time"`
}

func (s *SearchProfile) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchProfile() *SearchProfile { _ = "STUB: not implemented"; return nil }
