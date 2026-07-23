package catsnapshotscolumn

type CatSnapshotsColumn struct {
	Name string
}

var (
	Id = CatSnapshotsColumn{"id"}

	Repository = CatSnapshotsColumn{"repository"}

	Status = CatSnapshotsColumn{"status"}

	Startepoch = CatSnapshotsColumn{"start_epoch"}

	Starttime = CatSnapshotsColumn{"start_time"}

	Endepoch = CatSnapshotsColumn{"end_epoch"}

	Endtime = CatSnapshotsColumn{"end_time"}

	Duration = CatSnapshotsColumn{"duration"}

	Indices = CatSnapshotsColumn{"indices"}

	Successfulshards = CatSnapshotsColumn{"successful_shards"}

	Failedshards = CatSnapshotsColumn{"failed_shards"}

	Totalshards = CatSnapshotsColumn{"total_shards"}

	Reason = CatSnapshotsColumn{"reason"}
)

func (c CatSnapshotsColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatSnapshotsColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatSnapshotsColumn) String() string { _ = "STUB: not implemented"; return "" }
