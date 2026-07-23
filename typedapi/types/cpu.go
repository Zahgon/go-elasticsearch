package types

type Cpu struct {
	LoadAverage   map[string]Float64 `json:"load_average,omitempty"`
	Percent       *int               `json:"percent,omitempty"`
	Sys           Duration           `json:"sys,omitempty"`
	SysInMillis   *int64             `json:"sys_in_millis,omitempty"`
	Total         Duration           `json:"total,omitempty"`
	TotalInMillis *int64             `json:"total_in_millis,omitempty"`
	User          Duration           `json:"user,omitempty"`
	UserInMillis  *int64             `json:"user_in_millis,omitempty"`
}

func (s *Cpu) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCpu() *Cpu { _ = "STUB: not implemented"; return nil }
