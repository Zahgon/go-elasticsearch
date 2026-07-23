package catanomalydetectorcolumn

type CatAnomalyDetectorColumn struct {
	Name string
}

var (
	Assignmentexplanation = CatAnomalyDetectorColumn{"assignment_explanation"}

	Bucketscount = CatAnomalyDetectorColumn{"buckets.count"}

	Bucketstimeexpavg = CatAnomalyDetectorColumn{"buckets.time.exp_avg"}

	Bucketstimeexpavghour = CatAnomalyDetectorColumn{"buckets.time.exp_avg_hour"}

	Bucketstimemax = CatAnomalyDetectorColumn{"buckets.time.max"}

	Bucketstimemin = CatAnomalyDetectorColumn{"buckets.time.min"}

	Bucketstimetotal = CatAnomalyDetectorColumn{"buckets.time.total"}

	Databuckets = CatAnomalyDetectorColumn{"data.buckets"}

	Dataearliestrecord = CatAnomalyDetectorColumn{"data.earliest_record"}

	Dataemptybuckets = CatAnomalyDetectorColumn{"data.empty_buckets"}

	Datainputbytes = CatAnomalyDetectorColumn{"data.input_bytes"}

	Datainputfields = CatAnomalyDetectorColumn{"data.input_fields"}

	Datainputrecords = CatAnomalyDetectorColumn{"data.input_records"}

	Datainvaliddates = CatAnomalyDetectorColumn{"data.invalid_dates"}

	Datalast = CatAnomalyDetectorColumn{"data.last"}

	Datalastemptybucket = CatAnomalyDetectorColumn{"data.last_empty_bucket"}

	Datalastsparsebucket = CatAnomalyDetectorColumn{"data.last_sparse_bucket"}

	Datalatestrecord = CatAnomalyDetectorColumn{"data.latest_record"}

	Datamissingfields = CatAnomalyDetectorColumn{"data.missing_fields"}

	Dataoutofordertimestamps = CatAnomalyDetectorColumn{"data.out_of_order_timestamps"}

	Dataprocessedfields = CatAnomalyDetectorColumn{"data.processed_fields"}

	Dataprocessedrecords = CatAnomalyDetectorColumn{"data.processed_records"}

	Datasparsebuckets = CatAnomalyDetectorColumn{"data.sparse_buckets"}

	Forecastsmemoryavg = CatAnomalyDetectorColumn{"forecasts.memory.avg"}

	Forecastsmemorymax = CatAnomalyDetectorColumn{"forecasts.memory.max"}

	Forecastsmemorymin = CatAnomalyDetectorColumn{"forecasts.memory.min"}

	Forecastsmemorytotal = CatAnomalyDetectorColumn{"forecasts.memory.total"}

	Forecastsrecordsavg = CatAnomalyDetectorColumn{"forecasts.records.avg"}

	Forecastsrecordsmax = CatAnomalyDetectorColumn{"forecasts.records.max"}

	Forecastsrecordsmin = CatAnomalyDetectorColumn{"forecasts.records.min"}

	Forecastsrecordstotal = CatAnomalyDetectorColumn{"forecasts.records.total"}

	Forecaststimeavg = CatAnomalyDetectorColumn{"forecasts.time.avg"}

	Forecaststimemax = CatAnomalyDetectorColumn{"forecasts.time.max"}

	Forecaststimemin = CatAnomalyDetectorColumn{"forecasts.time.min"}

	Forecaststimetotal = CatAnomalyDetectorColumn{"forecasts.time.total"}

	Forecaststotal = CatAnomalyDetectorColumn{"forecasts.total"}

	Id = CatAnomalyDetectorColumn{"id"}

	Modelbucketallocationfailures = CatAnomalyDetectorColumn{"model.bucket_allocation_failures"}

	Modelbyfields = CatAnomalyDetectorColumn{"model.by_fields"}

	Modelbytes = CatAnomalyDetectorColumn{"model.bytes"}

	Modelbytesexceeded = CatAnomalyDetectorColumn{"model.bytes_exceeded"}

	Modelcategorizationstatus = CatAnomalyDetectorColumn{"model.categorization_status"}

	Modelcategorizeddoccount = CatAnomalyDetectorColumn{"model.categorized_doc_count"}

	Modeldeadcategorycount = CatAnomalyDetectorColumn{"model.dead_category_count"}

	Modelfailedcategorycount = CatAnomalyDetectorColumn{"model.failed_category_count"}

	Modelfrequentcategorycount = CatAnomalyDetectorColumn{"model.frequent_category_count"}

	Modellogtime = CatAnomalyDetectorColumn{"model.log_time"}

	Modelmemorylimit = CatAnomalyDetectorColumn{"model.memory_limit"}

	Modelmemorystatus = CatAnomalyDetectorColumn{"model.memory_status"}

	Modeloverfields = CatAnomalyDetectorColumn{"model.over_fields"}

	Modelpartitionfields = CatAnomalyDetectorColumn{"model.partition_fields"}

	Modelrarecategorycount = CatAnomalyDetectorColumn{"model.rare_category_count"}

	Modeltimestamp = CatAnomalyDetectorColumn{"model.timestamp"}

	Modeltotalcategorycount = CatAnomalyDetectorColumn{"model.total_category_count"}

	Nodeaddress = CatAnomalyDetectorColumn{"node.address"}

	Nodeephemeralid = CatAnomalyDetectorColumn{"node.ephemeral_id"}

	Nodeid = CatAnomalyDetectorColumn{"node.id"}

	Nodename = CatAnomalyDetectorColumn{"node.name"}

	Openedtime = CatAnomalyDetectorColumn{"opened_time"}

	State = CatAnomalyDetectorColumn{"state"}
)

func (c CatAnomalyDetectorColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatAnomalyDetectorColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatAnomalyDetectorColumn) String() string { _ = "STUB: not implemented"; return "" }
