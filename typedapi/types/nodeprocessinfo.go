package types

type NodeProcessInfo struct {
	Id int64 `json:"id"`

	Mlockall bool `json:"mlockall"`

	RefreshIntervalInMillis int64 `json:"refresh_interval_in_millis"`
}

func (s *NodeProcessInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeProcessInfo() *NodeProcessInfo { _ = "STUB: not implemented"; return nil }
