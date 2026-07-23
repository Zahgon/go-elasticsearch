package types

type ExecutionThreadPool struct {
	MaxSize int64 `json:"max_size"`

	QueueSize int64 `json:"queue_size"`
}

func (s *ExecutionThreadPool) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExecutionThreadPool() *ExecutionThreadPool { _ = "STUB: not implemented"; return nil }
