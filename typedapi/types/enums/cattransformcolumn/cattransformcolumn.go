package cattransformcolumn

type CatTransformColumn struct {
	Name string
}

var (
	Changeslastdetectiontime = CatTransformColumn{"changes_last_detection_time"}

	Checkpoint = CatTransformColumn{"checkpoint"}

	Checkpointdurationtimeexpavg = CatTransformColumn{"checkpoint_duration_time_exp_avg"}

	Checkpointprogress = CatTransformColumn{"checkpoint_progress"}

	Createtime = CatTransformColumn{"create_time"}

	Deletetime = CatTransformColumn{"delete_time"}

	Description = CatTransformColumn{"description"}

	Destindex = CatTransformColumn{"dest_index"}

	Documentsdeleted = CatTransformColumn{"documents_deleted"}

	Documentsindexed = CatTransformColumn{"documents_indexed"}

	Docspersecond = CatTransformColumn{"docs_per_second"}

	Documentsprocessed = CatTransformColumn{"documents_processed"}

	Frequency = CatTransformColumn{"frequency"}

	Id = CatTransformColumn{"id"}

	Indexfailure = CatTransformColumn{"index_failure"}

	Indextime = CatTransformColumn{"index_time"}

	Indextotal = CatTransformColumn{"index_total"}

	Indexeddocumentsexpavg = CatTransformColumn{"indexed_documents_exp_avg"}

	Lastsearchtime = CatTransformColumn{"last_search_time"}

	Maxpagesearchsize = CatTransformColumn{"max_page_search_size"}

	Pagesprocessed = CatTransformColumn{"pages_processed"}

	Pipeline = CatTransformColumn{"pipeline"}

	Processeddocumentsexpavg = CatTransformColumn{"processed_documents_exp_avg"}

	Processingtime = CatTransformColumn{"processing_time"}

	Projectrouting = CatTransformColumn{"project_routing"}

	Reason = CatTransformColumn{"reason"}

	Searchfailure = CatTransformColumn{"search_failure"}

	Searchtime = CatTransformColumn{"search_time"}

	Searchtotal = CatTransformColumn{"search_total"}

	Sourceindex = CatTransformColumn{"source_index"}

	State = CatTransformColumn{"state"}

	Transformtype = CatTransformColumn{"transform_type"}

	Triggercount = CatTransformColumn{"trigger_count"}

	Version = CatTransformColumn{"version"}
)

func (c CatTransformColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatTransformColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatTransformColumn) String() string { _ = "STUB: not implemented"; return "" }
