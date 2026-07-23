package types

type Collector struct {
	Children    []Collector `json:"children,omitempty"`
	Name        string      `json:"name"`
	Reason      string      `json:"reason"`
	TimeInNanos int64       `json:"time_in_nanos"`
}

func (s *Collector) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCollector() *Collector { _ = "STUB: not implemented"; return nil }
