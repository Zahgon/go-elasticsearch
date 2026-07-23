package types

type ReindexTask struct {
	Action             string              `json:"action"`
	Cancellable        bool                `json:"cancellable"`
	Cancelled          bool                `json:"cancelled"`
	Description        string              `json:"description"`
	Headers            HttpHeaders         `json:"headers"`
	Id                 int64               `json:"id"`
	Node               string              `json:"node"`
	RunningTimeInNanos int64               `json:"running_time_in_nanos"`
	StartTimeInMillis  int64               `json:"start_time_in_millis"`
	Status             ParentReindexStatus `json:"status"`
	Type               string              `json:"type"`
}

func (s *ReindexTask) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReindexTask() *ReindexTask { _ = "STUB: not implemented"; return nil }
