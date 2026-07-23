package types

type IpFilter struct {
	Http      bool `json:"http"`
	Transport bool `json:"transport"`
}

func (s *IpFilter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIpFilter() *IpFilter { _ = "STUB: not implemented"; return nil }
