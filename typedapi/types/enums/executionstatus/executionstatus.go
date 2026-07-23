package executionstatus

type ExecutionStatus struct {
	Name string
}

var (
	Awaitsexecution = ExecutionStatus{"awaits_execution"}

	Checking = ExecutionStatus{"checking"}

	Executionnotneeded = ExecutionStatus{"execution_not_needed"}

	Throttled = ExecutionStatus{"throttled"}

	Executed = ExecutionStatus{"executed"}

	Failed = ExecutionStatus{"failed"}

	Deletedwhilequeued = ExecutionStatus{"deleted_while_queued"}

	Notexecutedalreadyqueued = ExecutionStatus{"not_executed_already_queued"}
)

func (e ExecutionStatus) MarshalText() (text []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *ExecutionStatus) UnmarshalText(text []byte) error { _ = "STUB: not implemented"; return nil }

func (e ExecutionStatus) String() string { _ = "STUB: not implemented"; return "" }
