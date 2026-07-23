package cattaskscolumn

type CatTasksColumn struct {
	Name string
}

var (
	Id = CatTasksColumn{"id"}

	Action = CatTasksColumn{"action"}

	Taskid = CatTasksColumn{"task_id"}

	Parenttaskid = CatTasksColumn{"parent_task_id"}

	Type = CatTasksColumn{"type"}

	Starttime = CatTasksColumn{"start_time"}

	Timestamp = CatTasksColumn{"timestamp"}

	Runningtimens = CatTasksColumn{"running_time_ns"}

	Runningtime = CatTasksColumn{"running_time"}

	Nodeid = CatTasksColumn{"node_id"}

	Ip = CatTasksColumn{"ip"}

	Port = CatTasksColumn{"port"}

	Node = CatTasksColumn{"node"}

	Version = CatTasksColumn{"version"}

	Xopaqueid = CatTasksColumn{"x_opaque_id"}
)

func (c CatTasksColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatTasksColumn) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (c CatTasksColumn) String() string { _ = "STUB: not implemented"; return "" }
