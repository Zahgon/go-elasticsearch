package searchtype

type SearchType struct {
	Name string
}

var (
	Querythenfetch = SearchType{"query_then_fetch"}

	Dfsquerythenfetch = SearchType{"dfs_query_then_fetch"}
)

func (s SearchType) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *SearchType) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (s SearchType) String() string { _ = "STUB: not implemented"; return "" }
