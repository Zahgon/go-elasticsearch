package catshardcolumn

type CatShardColumn struct {
	Name string
}

var (
	Completionsize = CatShardColumn{"completion.size"}

	Datasetsize = CatShardColumn{"dataset.size"}

	Densevectorvaluecount = CatShardColumn{"dense_vector.value_count"}

	Docs = CatShardColumn{"docs"}

	Fielddataevictions = CatShardColumn{"fielddata.evictions"}

	Fielddatamemorysize = CatShardColumn{"fielddata.memory_size"}

	Flushtotal = CatShardColumn{"flush.total"}

	Flushtotaltime = CatShardColumn{"flush.total_time"}

	Getcurrent = CatShardColumn{"get.current"}

	Getexiststime = CatShardColumn{"get.exists_time"}

	Getexiststotal = CatShardColumn{"get.exists_total"}

	Getmissingtime = CatShardColumn{"get.missing_time"}

	Getmissingtotal = CatShardColumn{"get.missing_total"}

	Gettime = CatShardColumn{"get.time"}

	Gettotal = CatShardColumn{"get.total"}

	Id = CatShardColumn{"id"}

	Index = CatShardColumn{"index"}

	Indexingdeletecurrent = CatShardColumn{"indexing.delete_current"}

	Indexingdeletetime = CatShardColumn{"indexing.delete_time"}

	Indexingdeletetotal = CatShardColumn{"indexing.delete_total"}

	Indexingindexcurrent = CatShardColumn{"indexing.index_current"}

	Indexingindexfailedduetoversionconflict = CatShardColumn{"indexing.index_failed_due_to_version_conflict"}

	Indexingindexfailed = CatShardColumn{"indexing.index_failed"}

	Indexingindextime = CatShardColumn{"indexing.index_time"}

	Indexingindextotal = CatShardColumn{"indexing.index_total"}

	Ip = CatShardColumn{"ip"}

	Mergescurrent = CatShardColumn{"merges.current"}

	Mergescurrentdocs = CatShardColumn{"merges.current_docs"}

	Mergescurrentsize = CatShardColumn{"merges.current_size"}

	Mergestotal = CatShardColumn{"merges.total"}

	Mergestotaldocs = CatShardColumn{"merges.total_docs"}

	Mergestotalsize = CatShardColumn{"merges.total_size"}

	Mergestotaltime = CatShardColumn{"merges.total_time"}

	Node = CatShardColumn{"node"}

	Prirep = CatShardColumn{"prirep"}

	Querycacheevictions = CatShardColumn{"query_cache.evictions"}

	Querycachememorysize = CatShardColumn{"query_cache.memory_size"}

	Recoverysourcetype = CatShardColumn{"recoverysource.type"}

	Refreshtime = CatShardColumn{"refresh.time"}

	Refreshtotal = CatShardColumn{"refresh.total"}

	Searchfetchcurrent = CatShardColumn{"search.fetch_current"}

	Searchfetchtime = CatShardColumn{"search.fetch_time"}

	Searchfetchtotal = CatShardColumn{"search.fetch_total"}

	Searchopencontexts = CatShardColumn{"search.open_contexts"}

	Searchquerycurrent = CatShardColumn{"search.query_current"}

	Searchquerytime = CatShardColumn{"search.query_time"}

	Searchquerytotal = CatShardColumn{"search.query_total"}

	Searchscrollcurrent = CatShardColumn{"search.scroll_current"}

	Searchscrolltime = CatShardColumn{"search.scroll_time"}

	Searchscrolltotal = CatShardColumn{"search.scroll_total"}

	Segmentscount = CatShardColumn{"segments.count"}

	Segmentsfixedbitsetmemory = CatShardColumn{"segments.fixed_bitset_memory"}

	Segmentsindexwritermemory = CatShardColumn{"segments.index_writer_memory"}

	Segmentsmemory = CatShardColumn{"segments.memory"}

	Segmentsversionmapmemory = CatShardColumn{"segments.version_map_memory"}

	Seqnoglobalcheckpoint = CatShardColumn{"seq_no.global_checkpoint"}

	Seqnolocalcheckpoint = CatShardColumn{"seq_no.local_checkpoint"}

	Seqnomax = CatShardColumn{"seq_no.max"}

	Shard = CatShardColumn{"shard"}

	Dsparsevectorvaluecount = CatShardColumn{"dsparse_vector.value_count"}

	State = CatShardColumn{"state"}

	Store = CatShardColumn{"store"}

	Suggestcurrent = CatShardColumn{"suggest.current"}

	Suggesttime = CatShardColumn{"suggest.time"}

	Suggesttotal = CatShardColumn{"suggest.total"}

	Syncid = CatShardColumn{"sync_id"}

	Unassignedat = CatShardColumn{"unassigned.at"}

	Unassigneddetails = CatShardColumn{"unassigned.details"}

	Unassignedfor = CatShardColumn{"unassigned.for"}

	Unassignedreason = CatShardColumn{"unassigned.reason"}
)

func (c CatShardColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatShardColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatShardColumn) String() string { _ = "STUB: not implemented"; return "" }
