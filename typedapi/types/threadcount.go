package types

type ThreadCount struct {
	Active *int64 `json:"active,omitempty"`

	Completed *int64 `json:"completed,omitempty"`

	Largest *int64 `json:"largest,omitempty"`

	Queue *int64 `json:"queue,omitempty"`

	Rejected *int64 `json:"rejected,omitempty"`

	Threads *int64 `json:"threads,omitempty"`
}

func (s *ThreadCount) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewThreadCount() *ThreadCount { _ = "STUB: not implemented"; return nil }
