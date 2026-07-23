package catrecoverycolumn

type CatRecoveryColumn struct {
	Name string
}

var (
	Index = CatRecoveryColumn{"index"}

	Shard = CatRecoveryColumn{"shard"}

	Starttime = CatRecoveryColumn{"start_time"}

	Starttimemillis = CatRecoveryColumn{"start_time_millis"}

	Stoptime = CatRecoveryColumn{"stop_time"}

	Stoptimemillis = CatRecoveryColumn{"stop_time_millis"}

	Time = CatRecoveryColumn{"time"}

	Type = CatRecoveryColumn{"type"}

	Stage = CatRecoveryColumn{"stage"}

	Sourcehost = CatRecoveryColumn{"source_host"}

	Sourcenode = CatRecoveryColumn{"source_node"}

	Targethost = CatRecoveryColumn{"target_host"}

	Targetnode = CatRecoveryColumn{"target_node"}

	Repository = CatRecoveryColumn{"repository"}

	Snapshot = CatRecoveryColumn{"snapshot"}

	Files = CatRecoveryColumn{"files"}

	Filesrecovered = CatRecoveryColumn{"files_recovered"}

	Filespercent = CatRecoveryColumn{"files_percent"}

	Filestotal = CatRecoveryColumn{"files_total"}

	Bytes = CatRecoveryColumn{"bytes"}

	Bytesrecovered = CatRecoveryColumn{"bytes_recovered"}

	Bytespercent = CatRecoveryColumn{"bytes_percent"}

	Bytestotal = CatRecoveryColumn{"bytes_total"}

	Translogops = CatRecoveryColumn{"translog_ops"}

	Translogopsrecovered = CatRecoveryColumn{"translog_ops_recovered"}

	Translogopspercent = CatRecoveryColumn{"translog_ops_percent"}
)

func (c CatRecoveryColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatRecoveryColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatRecoveryColumn) String() string { _ = "STUB: not implemented"; return "" }
