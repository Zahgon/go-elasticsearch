package catthreadpoolcolumn

type CatThreadPoolColumn struct {
	Name string
}

var (
	Active = CatThreadPoolColumn{"active"}

	Completed = CatThreadPoolColumn{"completed"}

	Core = CatThreadPoolColumn{"core"}

	Ephemeralid = CatThreadPoolColumn{"ephemeral_id"}

	Host = CatThreadPoolColumn{"host"}

	Ip = CatThreadPoolColumn{"ip"}

	Keepalive = CatThreadPoolColumn{"keep_alive"}

	Largest = CatThreadPoolColumn{"largest"}

	Max = CatThreadPoolColumn{"max"}

	Name = CatThreadPoolColumn{"name"}

	Nodeid = CatThreadPoolColumn{"node_id"}

	Nodename = CatThreadPoolColumn{"node_name"}

	Pid = CatThreadPoolColumn{"pid"}

	Poolsize = CatThreadPoolColumn{"pool_size"}

	Port = CatThreadPoolColumn{"port"}

	Queue = CatThreadPoolColumn{"queue"}

	Queuesize = CatThreadPoolColumn{"queue_size"}

	Rejected = CatThreadPoolColumn{"rejected"}

	Size = CatThreadPoolColumn{"size"}

	Type = CatThreadPoolColumn{"type"}
)

func (c CatThreadPoolColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatThreadPoolColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatThreadPoolColumn) String() string { _ = "STUB: not implemented"; return "" }
