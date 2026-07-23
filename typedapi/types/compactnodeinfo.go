package types

type CompactNodeInfo struct {
	Name string `json:"name"`
}

func (s *CompactNodeInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompactNodeInfo() *CompactNodeInfo { _ = "STUB: not implemented"; return nil }
