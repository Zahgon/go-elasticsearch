package types

type AllocationRecord struct {
	DiskAvail ByteSize `json:"disk.avail,omitempty"`

	DiskIndices ByteSize `json:"disk.indices,omitempty"`

	DiskIndicesForecast ByteSize `json:"disk.indices.forecast,omitempty"`

	DiskPercent Percentage `json:"disk.percent,omitempty"`

	DiskTotal ByteSize `json:"disk.total,omitempty"`

	DiskUsed ByteSize `json:"disk.used,omitempty"`

	Host *string `json:"host,omitempty"`

	Ip *string `json:"ip,omitempty"`

	Node *string `json:"node,omitempty"`

	NodeRole *string `json:"node.role,omitempty"`

	Shards *string `json:"shards,omitempty"`

	ShardsUndesired *string `json:"shards.undesired,omitempty"`

	WriteLoadForecast Stringifieddouble `json:"write_load.forecast,omitempty"`
}

func (s *AllocationRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAllocationRecord() *AllocationRecord { _ = "STUB: not implemented"; return nil }
