package types

type SearchStats struct {
	FetchCurrent        int64                  `json:"fetch_current"`
	FetchTime           Duration               `json:"fetch_time,omitempty"`
	FetchTimeInMillis   int64                  `json:"fetch_time_in_millis"`
	FetchTotal          int64                  `json:"fetch_total"`
	Groups              map[string]SearchStats `json:"groups,omitempty"`
	OpenContexts        *int64                 `json:"open_contexts,omitempty"`
	QueryCurrent        int64                  `json:"query_current"`
	QueryTime           Duration               `json:"query_time,omitempty"`
	QueryTimeInMillis   int64                  `json:"query_time_in_millis"`
	QueryTotal          int64                  `json:"query_total"`
	RecentSearchLoad    *Float64               `json:"recent_search_load,omitempty"`
	ScrollCurrent       int64                  `json:"scroll_current"`
	ScrollTime          Duration               `json:"scroll_time,omitempty"`
	ScrollTimeInMillis  int64                  `json:"scroll_time_in_millis"`
	ScrollTotal         int64                  `json:"scroll_total"`
	SuggestCurrent      int64                  `json:"suggest_current"`
	SuggestTime         Duration               `json:"suggest_time,omitempty"`
	SuggestTimeInMillis int64                  `json:"suggest_time_in_millis"`
	SuggestTotal        int64                  `json:"suggest_total"`
}

func (s *SearchStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSearchStats() *SearchStats { _ = "STUB: not implemented"; return nil }
