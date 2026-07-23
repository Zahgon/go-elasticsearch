package cattrainedmodelscolumn

type CatTrainedModelsColumn struct {
	Name string
}

var (
	Createtime = CatTrainedModelsColumn{"create_time"}

	Createdby = CatTrainedModelsColumn{"created_by"}

	Dataframeanalyticsid = CatTrainedModelsColumn{"data_frame_analytics_id"}

	Description = CatTrainedModelsColumn{"description"}

	Heapsize = CatTrainedModelsColumn{"heap_size"}

	Id = CatTrainedModelsColumn{"id"}

	Ingestcount = CatTrainedModelsColumn{"ingest.count"}

	Ingestcurrent = CatTrainedModelsColumn{"ingest.current"}

	Ingestfailed = CatTrainedModelsColumn{"ingest.failed"}

	Ingestpipelines = CatTrainedModelsColumn{"ingest.pipelines"}

	Ingesttime = CatTrainedModelsColumn{"ingest.time"}

	License = CatTrainedModelsColumn{"license"}

	Operations = CatTrainedModelsColumn{"operations"}

	Version = CatTrainedModelsColumn{"version"}
)

func (c CatTrainedModelsColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatTrainedModelsColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatTrainedModelsColumn) String() string { _ = "STUB: not implemented"; return "" }
