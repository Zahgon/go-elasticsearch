package catindicescolumn

type CatIndicesColumn struct {
	Name string
}

var (
	Health = CatIndicesColumn{"health"}

	Status = CatIndicesColumn{"status"}

	Index = CatIndicesColumn{"index"}

	Uuid = CatIndicesColumn{"uuid"}

	Pri = CatIndicesColumn{"pri"}

	Rep = CatIndicesColumn{"rep"}

	Docscount = CatIndicesColumn{"docs.count"}

	Docsdeleted = CatIndicesColumn{"docs.deleted"}

	Creationdate = CatIndicesColumn{"creation.date"}

	Creationdatestring = CatIndicesColumn{"creation.date.string"}

	Storesize = CatIndicesColumn{"store.size"}

	Pristoresize = CatIndicesColumn{"pri.store.size"}

	Datasetsize = CatIndicesColumn{"dataset.size"}

	Completionsize = CatIndicesColumn{"completion.size"}

	Pricompletionsize = CatIndicesColumn{"pri.completion.size"}

	Fielddatamemorysize = CatIndicesColumn{"fielddata.memory_size"}

	Prifielddatamemorysize = CatIndicesColumn{"pri.fielddata.memory_size"}

	Fielddataevictions = CatIndicesColumn{"fielddata.evictions"}

	Prifielddataevictions = CatIndicesColumn{"pri.fielddata.evictions"}

	Querycachememorysize = CatIndicesColumn{"query_cache.memory_size"}

	Priquerycachememorysize = CatIndicesColumn{"pri.query_cache.memory_size"}

	Querycacheevictions = CatIndicesColumn{"query_cache.evictions"}

	Priquerycacheevictions = CatIndicesColumn{"pri.query_cache.evictions"}

	Requestcachememorysize = CatIndicesColumn{"request_cache.memory_size"}

	Prirequestcachememorysize = CatIndicesColumn{"pri.request_cache.memory_size"}

	Requestcacheevictions = CatIndicesColumn{"request_cache.evictions"}

	Prirequestcacheevictions = CatIndicesColumn{"pri.request_cache.evictions"}

	Requestcachehitcount = CatIndicesColumn{"request_cache.hit_count"}

	Prirequestcachehitcount = CatIndicesColumn{"pri.request_cache.hit_count"}

	Requestcachemisscount = CatIndicesColumn{"request_cache.miss_count"}

	Prirequestcachemisscount = CatIndicesColumn{"pri.request_cache.miss_count"}

	Flushtotal = CatIndicesColumn{"flush.total"}

	Priflushtotal = CatIndicesColumn{"pri.flush.total"}

	Flushtotaltime = CatIndicesColumn{"flush.total_time"}

	Priflushtotaltime = CatIndicesColumn{"pri.flush.total_time"}

	Getcurrent = CatIndicesColumn{"get.current"}

	Prigetcurrent = CatIndicesColumn{"pri.get.current"}

	Gettime = CatIndicesColumn{"get.time"}

	Prigettime = CatIndicesColumn{"pri.get.time"}

	Gettotal = CatIndicesColumn{"get.total"}

	Prigettotal = CatIndicesColumn{"pri.get.total"}

	Getexiststime = CatIndicesColumn{"get.exists_time"}

	Prigetexiststime = CatIndicesColumn{"pri.get.exists_time"}

	Getexiststotal = CatIndicesColumn{"get.exists_total"}

	Prigetexiststotal = CatIndicesColumn{"pri.get.exists_total"}

	Getmissingtime = CatIndicesColumn{"get.missing_time"}

	Prigetmissingtime = CatIndicesColumn{"pri.get.missing_time"}

	Getmissingtotal = CatIndicesColumn{"get.missing_total"}

	Prigetmissingtotal = CatIndicesColumn{"pri.get.missing_total"}

	Indexingdeletecurrent = CatIndicesColumn{"indexing.delete_current"}

	Priindexingdeletecurrent = CatIndicesColumn{"pri.indexing.delete_current"}

	Indexingdeletetime = CatIndicesColumn{"indexing.delete_time"}

	Priindexingdeletetime = CatIndicesColumn{"pri.indexing.delete_time"}

	Indexingdeletetotal = CatIndicesColumn{"indexing.delete_total"}

	Priindexingdeletetotal = CatIndicesColumn{"pri.indexing.delete_total"}

	Indexingindexcurrent = CatIndicesColumn{"indexing.index_current"}

	Priindexingindexcurrent = CatIndicesColumn{"pri.indexing.index_current"}

	Indexingindextime = CatIndicesColumn{"indexing.index_time"}

	Priindexingindextime = CatIndicesColumn{"pri.indexing.index_time"}

	Indexingindextotal = CatIndicesColumn{"indexing.index_total"}

	Priindexingindextotal = CatIndicesColumn{"pri.indexing.index_total"}

	Indexingindexfailed = CatIndicesColumn{"indexing.index_failed"}

	Priindexingindexfailed = CatIndicesColumn{"pri.indexing.index_failed"}

	Indexingindexfailedduetoversionconflict = CatIndicesColumn{"indexing.index_failed_due_to_version_conflict"}

	Priindexingindexfailedduetoversionconflict = CatIndicesColumn{"pri.indexing.index_failed_due_to_version_conflict"}

	Mergescurrent = CatIndicesColumn{"merges.current"}

	Primergescurrent = CatIndicesColumn{"pri.merges.current"}

	Mergescurrentdocs = CatIndicesColumn{"merges.current_docs"}

	Primergescurrentdocs = CatIndicesColumn{"pri.merges.current_docs"}

	Mergescurrentsize = CatIndicesColumn{"merges.current_size"}

	Primergescurrentsize = CatIndicesColumn{"pri.merges.current_size"}

	Mergestotal = CatIndicesColumn{"merges.total"}

	Primergestotal = CatIndicesColumn{"pri.merges.total"}

	Mergestotaldocs = CatIndicesColumn{"merges.total_docs"}

	Primergestotaldocs = CatIndicesColumn{"pri.merges.total_docs"}

	Mergestotalsize = CatIndicesColumn{"merges.total_size"}

	Primergestotalsize = CatIndicesColumn{"pri.merges.total_size"}

	Mergestotaltime = CatIndicesColumn{"merges.total_time"}

	Primergestotaltime = CatIndicesColumn{"pri.merges.total_time"}

	Refreshtotal = CatIndicesColumn{"refresh.total"}

	Prirefreshtotal = CatIndicesColumn{"pri.refresh.total"}

	Refreshtime = CatIndicesColumn{"refresh.time"}

	Prirefreshtime = CatIndicesColumn{"pri.refresh.time"}

	Refreshexternaltotal = CatIndicesColumn{"refresh.external_total"}

	Prirefreshexternaltotal = CatIndicesColumn{"pri.refresh.external_total"}

	Refreshexternaltime = CatIndicesColumn{"refresh.external_time"}

	Prirefreshexternaltime = CatIndicesColumn{"pri.refresh.external_time"}

	Refreshlisteners = CatIndicesColumn{"refresh.listeners"}

	Prirefreshlisteners = CatIndicesColumn{"pri.refresh.listeners"}

	Searchfetchcurrent = CatIndicesColumn{"search.fetch_current"}

	Prisearchfetchcurrent = CatIndicesColumn{"pri.search.fetch_current"}

	Searchfetchtime = CatIndicesColumn{"search.fetch_time"}

	Prisearchfetchtime = CatIndicesColumn{"pri.search.fetch_time"}

	Searchfetchtotal = CatIndicesColumn{"search.fetch_total"}

	Prisearchfetchtotal = CatIndicesColumn{"pri.search.fetch_total"}

	Searchopencontexts = CatIndicesColumn{"search.open_contexts"}

	Prisearchopencontexts = CatIndicesColumn{"pri.search.open_contexts"}

	Searchquerycurrent = CatIndicesColumn{"search.query_current"}

	Prisearchquerycurrent = CatIndicesColumn{"pri.search.query_current"}

	Searchquerytime = CatIndicesColumn{"search.query_time"}

	Prisearchquerytime = CatIndicesColumn{"pri.search.query_time"}

	Searchquerytotal = CatIndicesColumn{"search.query_total"}

	Prisearchquerytotal = CatIndicesColumn{"pri.search.query_total"}

	Searchscrollcurrent = CatIndicesColumn{"search.scroll_current"}

	Prisearchscrollcurrent = CatIndicesColumn{"pri.search.scroll_current"}

	Searchscrolltime = CatIndicesColumn{"search.scroll_time"}

	Prisearchscrolltime = CatIndicesColumn{"pri.search.scroll_time"}

	Searchscrolltotal = CatIndicesColumn{"search.scroll_total"}

	Prisearchscrolltotal = CatIndicesColumn{"pri.search.scroll_total"}

	Segmentscount = CatIndicesColumn{"segments.count"}

	Prisegmentscount = CatIndicesColumn{"pri.segments.count"}

	Segmentsmemory = CatIndicesColumn{"segments.memory"}

	Prisegmentsmemory = CatIndicesColumn{"pri.segments.memory"}

	Segmentsindexwritermemory = CatIndicesColumn{"segments.index_writer_memory"}

	Prisegmentsindexwritermemory = CatIndicesColumn{"pri.segments.index_writer_memory"}

	Segmentsversionmapmemory = CatIndicesColumn{"segments.version_map_memory"}

	Prisegmentsversionmapmemory = CatIndicesColumn{"pri.segments.version_map_memory"}

	Segmentsfixedbitsetmemory = CatIndicesColumn{"segments.fixed_bitset_memory"}

	Prisegmentsfixedbitsetmemory = CatIndicesColumn{"pri.segments.fixed_bitset_memory"}

	Warmercurrent = CatIndicesColumn{"warmer.current"}

	Priwarmercurrent = CatIndicesColumn{"pri.warmer.current"}

	Warmertotal = CatIndicesColumn{"warmer.total"}

	Priwarmertotal = CatIndicesColumn{"pri.warmer.total"}

	Warmertotaltime = CatIndicesColumn{"warmer.total_time"}

	Priwarmertotaltime = CatIndicesColumn{"pri.warmer.total_time"}

	Suggestcurrent = CatIndicesColumn{"suggest.current"}

	Prisuggestcurrent = CatIndicesColumn{"pri.suggest.current"}

	Suggesttime = CatIndicesColumn{"suggest.time"}

	Prisuggesttime = CatIndicesColumn{"pri.suggest.time"}

	Suggesttotal = CatIndicesColumn{"suggest.total"}

	Prisuggesttotal = CatIndicesColumn{"pri.suggest.total"}

	Memorytotal = CatIndicesColumn{"memory.total"}

	Primemorytotal = CatIndicesColumn{"pri.memory.total"}

	Bulktotaloperations = CatIndicesColumn{"bulk.total_operations"}

	Pribulktotaloperations = CatIndicesColumn{"pri.bulk.total_operations"}

	Bulktotaltime = CatIndicesColumn{"bulk.total_time"}

	Pribulktotaltime = CatIndicesColumn{"pri.bulk.total_time"}

	Bulktotalsizeinbytes = CatIndicesColumn{"bulk.total_size_in_bytes"}

	Pribulktotalsizeinbytes = CatIndicesColumn{"pri.bulk.total_size_in_bytes"}

	Bulkavgtime = CatIndicesColumn{"bulk.avg_time"}

	Pribulkavgtime = CatIndicesColumn{"pri.bulk.avg_time"}

	Bulkavgsizeinbytes = CatIndicesColumn{"bulk.avg_size_in_bytes"}

	Pribulkavgsizeinbytes = CatIndicesColumn{"pri.bulk.avg_size_in_bytes"}

	Densevectorvaluecount = CatIndicesColumn{"dense_vector.value_count"}

	Pridensevectorvaluecount = CatIndicesColumn{"pri.dense_vector.value_count"}

	Sparsevectorvaluecount = CatIndicesColumn{"sparse_vector.value_count"}

	Prisparsevectorvaluecount = CatIndicesColumn{"pri.sparse_vector.value_count"}
)

func (c CatIndicesColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatIndicesColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatIndicesColumn) String() string { _ = "STUB: not implemented"; return "" }
