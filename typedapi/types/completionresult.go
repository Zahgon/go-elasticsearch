package types

type CompletionResult struct {
	Result string `json:"result"`
}

func (s *CompletionResult) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompletionResult() *CompletionResult { _ = "STUB: not implemented"; return nil }
