package cathealthcolumn

type CatHealthColumn struct {
	Name string
}

var (
	Epoch = CatHealthColumn{"epoch"}

	Timestamp = CatHealthColumn{"timestamp"}

	Cluster = CatHealthColumn{"cluster"}

	Status = CatHealthColumn{"status"}

	Nodetotal = CatHealthColumn{"node.total"}

	Nodedata = CatHealthColumn{"node.data"}

	Shards = CatHealthColumn{"shards"}

	Pri = CatHealthColumn{"pri"}

	Relo = CatHealthColumn{"relo"}

	Init = CatHealthColumn{"init"}

	Unassign = CatHealthColumn{"unassign"}

	Unassignpri = CatHealthColumn{"unassign.pri"}

	Pendingtasks = CatHealthColumn{"pending_tasks"}

	Maxtaskwaittime = CatHealthColumn{"max_task_wait_time"}

	Activeshardspercent = CatHealthColumn{"active_shards_percent"}
)

func (c CatHealthColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatHealthColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatHealthColumn) String() string { _ = "STUB: not implemented"; return "" }
