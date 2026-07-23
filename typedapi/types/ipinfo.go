package types

type Ipinfo struct {
}

func NewIpinfo() *Ipinfo { _ = "STUB: not implemented"; return nil }

type IpinfoVariant interface {
	IpinfoCaster() *Ipinfo
}

func (s *Ipinfo) IpinfoCaster() *Ipinfo { _ = "STUB: not implemented"; return nil }
