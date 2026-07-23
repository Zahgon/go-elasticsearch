package catnodecolumn

type CatNodeColumn struct {
	Name string
}

var (
	Build = CatNodeColumn{"build"}

	Completionsize = CatNodeColumn{"completion.size"}

	Cpu = CatNodeColumn{"cpu"}

	Diskavail = CatNodeColumn{"disk.avail"}

	Disktotal = CatNodeColumn{"disk.total"}

	Diskused = CatNodeColumn{"disk.used"}

	Diskusedpercent = CatNodeColumn{"disk.used_percent"}

	Fielddataevictions = CatNodeColumn{"fielddata.evictions"}

	Fielddatamemorysize = CatNodeColumn{"fielddata.memory_size"}

	Filedesccurrent = CatNodeColumn{"file_desc.current"}

	Filedescmax = CatNodeColumn{"file_desc.max"}

	Filedescpercent = CatNodeColumn{"file_desc.percent"}

	Flushtotal = CatNodeColumn{"flush.total"}

	Flushtotaltime = CatNodeColumn{"flush.total_time"}

	Getcurrent = CatNodeColumn{"get.current"}

	Getexiststime = CatNodeColumn{"get.exists_time"}

	Getexiststotal = CatNodeColumn{"get.exists_total"}

	Getmissingtime = CatNodeColumn{"get.missing_time"}

	Getmissingtotal = CatNodeColumn{"get.missing_total"}

	Gettime = CatNodeColumn{"get.time"}

	Gettotal = CatNodeColumn{"get.total"}

	Heapcurrent = CatNodeColumn{"heap.current"}

	Heapmax = CatNodeColumn{"heap.max"}

	Heappercent = CatNodeColumn{"heap.percent"}

	Httpaddress = CatNodeColumn{"http_address"}

	Id = CatNodeColumn{"id"}

	Indexingdeletecurrent = CatNodeColumn{"indexing.delete_current"}

	Indexingdeletetime = CatNodeColumn{"indexing.delete_time"}

	Indexingdeletetotal = CatNodeColumn{"indexing.delete_total"}

	Indexingindexcurrent = CatNodeColumn{"indexing.index_current"}

	Indexingindexfailed = CatNodeColumn{"indexing.index_failed"}

	Indexingindexfailedduetoversionconflict = CatNodeColumn{"indexing.index_failed_due_to_version_conflict"}

	Indexingindextime = CatNodeColumn{"indexing.index_time"}

	Indexingindextotal = CatNodeColumn{"indexing.index_total"}

	Ip = CatNodeColumn{"ip"}

	Jdk = CatNodeColumn{"jdk"}

	Load1m = CatNodeColumn{"load_1m"}

	Load5m = CatNodeColumn{"load_5m"}

	Load15m = CatNodeColumn{"load_15m"}

	Availableprocessors = CatNodeColumn{"available_processors"}

	Mappingstotalcount = CatNodeColumn{"mappings.total_count"}

	Mappingstotalestimatedoverheadinbytes = CatNodeColumn{"mappings.total_estimated_overhead_in_bytes"}

	Master = CatNodeColumn{"master"}

	Mergescurrent = CatNodeColumn{"merges.current"}

	Mergescurrentdocs = CatNodeColumn{"merges.current_docs"}

	Mergescurrentsize = CatNodeColumn{"merges.current_size"}

	Mergestotal = CatNodeColumn{"merges.total"}

	Mergestotaldocs = CatNodeColumn{"merges.total_docs"}

	Mergestotalsize = CatNodeColumn{"merges.total_size"}

	Mergestotaltime = CatNodeColumn{"merges.total_time"}

	Name = CatNodeColumn{"name"}

	Noderole = CatNodeColumn{"node.role"}

	Pid = CatNodeColumn{"pid"}

	Port = CatNodeColumn{"port"}

	Querycachememorysize = CatNodeColumn{"query_cache.memory_size"}

	Querycacheevictions = CatNodeColumn{"query_cache.evictions"}

	Querycachehitcount = CatNodeColumn{"query_cache.hit_count"}

	Querycachemisscount = CatNodeColumn{"query_cache.miss_count"}

	Ramcurrent = CatNodeColumn{"ram.current"}

	Rammax = CatNodeColumn{"ram.max"}

	Rampercent = CatNodeColumn{"ram.percent"}

	Refreshtotal = CatNodeColumn{"refresh.total"}

	Refreshtime = CatNodeColumn{"refresh.time"}

	Requestcachememorysize = CatNodeColumn{"request_cache.memory_size"}

	Requestcacheevictions = CatNodeColumn{"request_cache.evictions"}

	Requestcachehitcount = CatNodeColumn{"request_cache.hit_count"}

	Requestcachemisscount = CatNodeColumn{"request_cache.miss_count"}

	Scriptcompilations = CatNodeColumn{"script.compilations"}

	Scriptcacheevictions = CatNodeColumn{"script.cache_evictions"}

	Searchfetchcurrent = CatNodeColumn{"search.fetch_current"}

	Searchfetchtime = CatNodeColumn{"search.fetch_time"}

	Searchfetchtotal = CatNodeColumn{"search.fetch_total"}

	Searchopencontexts = CatNodeColumn{"search.open_contexts"}

	Searchquerycurrent = CatNodeColumn{"search.query_current"}

	Searchquerytime = CatNodeColumn{"search.query_time"}

	Searchquerytotal = CatNodeColumn{"search.query_total"}

	Searchscrollcurrent = CatNodeColumn{"search.scroll_current"}

	Searchscrolltime = CatNodeColumn{"search.scroll_time"}

	Searchscrolltotal = CatNodeColumn{"search.scroll_total"}

	Segmentscount = CatNodeColumn{"segments.count"}

	Segmentsfixedbitsetmemory = CatNodeColumn{"segments.fixed_bitset_memory"}

	Segmentsindexwritermemory = CatNodeColumn{"segments.index_writer_memory"}

	Segmentsmemory = CatNodeColumn{"segments.memory"}

	Segmentsversionmapmemory = CatNodeColumn{"segments.version_map_memory"}

	Shardstatstotalcount = CatNodeColumn{"shard_stats.total_count"}

	Suggestcurrent = CatNodeColumn{"suggest.current"}

	Suggesttime = CatNodeColumn{"suggest.time"}

	Suggesttotal = CatNodeColumn{"suggest.total"}

	Uptime = CatNodeColumn{"uptime"}

	Version = CatNodeColumn{"version"}
)

func (c CatNodeColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatNodeColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatNodeColumn) String() string { _ = "STUB: not implemented"; return "" }
