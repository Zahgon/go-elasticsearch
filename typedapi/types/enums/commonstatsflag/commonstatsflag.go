package commonstatsflag

type CommonStatsFlag struct {
	Name string
}

var (
	All = CommonStatsFlag{"_all"}

	Store = CommonStatsFlag{"store"}

	Indexing = CommonStatsFlag{"indexing"}

	Get = CommonStatsFlag{"get"}

	Search = CommonStatsFlag{"search"}

	Merge = CommonStatsFlag{"merge"}

	Flush = CommonStatsFlag{"flush"}

	Refresh = CommonStatsFlag{"refresh"}

	Querycache = CommonStatsFlag{"query_cache"}

	Fielddata = CommonStatsFlag{"fielddata"}

	Docs = CommonStatsFlag{"docs"}

	Warmer = CommonStatsFlag{"warmer"}

	Completion = CommonStatsFlag{"completion"}

	Segments = CommonStatsFlag{"segments"}

	Translog = CommonStatsFlag{"translog"}

	Requestcache = CommonStatsFlag{"request_cache"}

	Recovery = CommonStatsFlag{"recovery"}

	Bulk = CommonStatsFlag{"bulk"}

	Shardstats = CommonStatsFlag{"shard_stats"}

	Mappings = CommonStatsFlag{"mappings"}

	Densevector = CommonStatsFlag{"dense_vector"}

	Sparsevector = CommonStatsFlag{"sparse_vector"}
)

func (c CommonStatsFlag) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CommonStatsFlag) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CommonStatsFlag) String() string { _ = "STUB: not implemented"; return "" }
