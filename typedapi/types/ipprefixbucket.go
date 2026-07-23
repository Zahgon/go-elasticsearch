package types

type IpPrefixBucket struct {
	Aggregations map[string]Aggregate `json:"-"`
	DocCount     int64                `json:"doc_count"`
	IsIpv6       bool                 `json:"is_ipv6"`
	Key          string               `json:"key"`
	Netmask      *string              `json:"netmask,omitempty"`
	PrefixLength int                  `json:"prefix_length"`
}

func (s *IpPrefixBucket) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s IpPrefixBucket) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIpPrefixBucket() *IpPrefixBucket { _ = "STUB: not implemented"; return nil }
