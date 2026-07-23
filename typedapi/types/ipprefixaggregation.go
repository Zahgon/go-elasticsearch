package types

type IpPrefixAggregation struct {
	AppendPrefixLength *bool `json:"append_prefix_length,omitempty"`

	Field string `json:"field"`

	IsIpv6 *bool `json:"is_ipv6,omitempty"`

	Keyed *bool `json:"keyed,omitempty"`

	MinDocCount *int64 `json:"min_doc_count,omitempty"`

	PrefixLength int `json:"prefix_length"`
}

func (s *IpPrefixAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIpPrefixAggregation() *IpPrefixAggregation { _ = "STUB: not implemented"; return nil }

type IpPrefixAggregationVariant interface {
	IpPrefixAggregationCaster() *IpPrefixAggregation
}

func (s *IpPrefixAggregation) IpPrefixAggregationCaster() *IpPrefixAggregation {
	_ = "STUB: not implemented"
	return nil
}
