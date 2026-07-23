package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsPartition struct {
	v *types.TermsPartition
}

func NewTermsPartition(numpartitions int64, partition int64) *_termsPartition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsPartition) NumPartitions(numpartitions int64) *_termsPartition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsPartition) Partition(partition int64) *_termsPartition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsPartition) TermsPartitionCaster() *types.TermsPartition {
	_ = "STUB: not implemented"
	return nil
}
