package catpendingtaskscolumn

type CatPendingTasksColumn struct {
	Name string
}

var (
	InsertOrder = CatPendingTasksColumn{"insertOrder"}

	TimeInQueue = CatPendingTasksColumn{"timeInQueue"}

	Priority = CatPendingTasksColumn{"priority"}

	Source = CatPendingTasksColumn{"source"}
)

func (c CatPendingTasksColumn) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *CatPendingTasksColumn) UnmarshalText(text []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c CatPendingTasksColumn) String() string { _ = "STUB: not implemented"; return "" }
