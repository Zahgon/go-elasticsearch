package types

type ReindexTaskInfo struct {
	Cancelled bool `json:"cancelled"`

	Description *string `json:"description,omitempty"`

	Id string `json:"id"`

	RunningTime Duration `json:"running_time,omitempty"`

	RunningTimeInNanos int64 `json:"running_time_in_nanos"`

	StartTime *string `json:"start_time,omitempty"`

	StartTimeInMillis int64 `json:"start_time_in_millis"`

	Status *ReindexStatus `json:"status,omitempty"`
}

func (s *ReindexTaskInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReindexTaskInfo() *ReindexTaskInfo { _ = "STUB: not implemented"; return nil }
