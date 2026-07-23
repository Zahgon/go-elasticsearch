package types

type Body struct {
	Id               int64  `json:"id"`
	Node             string `json:"node"`
	Query            string `json:"query"`
	RunningTimeNanos int64  `json:"running_time_nanos"`
	StartTimeMillis  int64  `json:"start_time_millis"`
}

func (s *Body) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBody() *Body { _ = "STUB: not implemented"; return nil }
