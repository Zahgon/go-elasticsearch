package types

type InProgress struct {
	Name            string `json:"name"`
	StartTimeMillis int64  `json:"start_time_millis"`
	State           string `json:"state"`
	Uuid            string `json:"uuid"`
}

func (s *InProgress) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInProgress() *InProgress { _ = "STUB: not implemented"; return nil }
