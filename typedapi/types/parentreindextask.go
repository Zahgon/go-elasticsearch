package types

type ParentReindexTask struct {
	Action             string              `json:"action"`
	Cancellable        bool                `json:"cancellable"`
	Cancelled          bool                `json:"cancelled"`
	Children           []ReindexTask       `json:"children,omitempty"`
	Description        string              `json:"description"`
	Headers            HttpHeaders         `json:"headers"`
	Id                 int64               `json:"id"`
	Node               string              `json:"node"`
	RunningTimeInNanos int64               `json:"running_time_in_nanos"`
	StartTimeInMillis  int64               `json:"start_time_in_millis"`
	Status             ParentReindexStatus `json:"status"`
	Type               string              `json:"type"`
}

func (s *ParentReindexTask) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewParentReindexTask() *ParentReindexTask { _ = "STUB: not implemented"; return nil }
