package types

type ExecutionResult struct {
	Actions           []ExecutionResultAction  `json:"actions"`
	Condition         ExecutionResultCondition `json:"condition"`
	ExecutionDuration int64                    `json:"execution_duration"`
	ExecutionTime     DateTime                 `json:"execution_time"`
	Input             ExecutionResultInput     `json:"input"`
}

func (s *ExecutionResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewExecutionResult() *ExecutionResult { _ = "STUB: not implemented"; return nil }
