package types

type SearchUsageStats struct {
	Extended   ExtendedSearchUsage `json:"extended"`
	Queries    map[string]int64    `json:"queries"`
	Rescorers  map[string]int64    `json:"rescorers"`
	Retrievers map[string]int64    `json:"retrievers"`
	Sections   map[string]int64    `json:"sections"`
	Total      int64               `json:"total"`
}

func (s *SearchUsageStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchUsageStats() *SearchUsageStats { _ = "STUB: not implemented"; return nil }
