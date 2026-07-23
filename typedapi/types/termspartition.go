package types

type TermsPartition struct {
	NumPartitions int64 `json:"num_partitions"`

	Partition int64 `json:"partition"`
}

func (s *TermsPartition) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermsPartition() *TermsPartition { _ = "STUB: not implemented"; return nil }

type TermsPartitionVariant interface {
	TermsPartitionCaster() *TermsPartition
}

func (s *TermsPartition) TermsPartitionCaster() *TermsPartition {
	_ = "STUB: not implemented"
	return nil
}

func (s *TermsPartition) TermsIncludeCaster() *TermsInclude { _ = "STUB: not implemented"; return nil }
